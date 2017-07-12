package main

import (
	"errors"
	"io"
	"chess/srv/srv-room/client_handler"
	"chess/common/helper"
	"chess/common/log"
	pb "chess/srv/srv-room/proto"
	"google.golang.org/grpc/metadata"
	"strconv"
	"chess/srv/srv-room/registry"
	. "chess/srv/srv-room/types"
	"encoding/binary"
)

const (
	DEFAULT_CH_IPC_SIZE = 16 // 默认玩家异步IPC消息队列大小
)

var (
	ERROR_INCORRECT_FRAME_TYPE = errors.New("incorrect frame type")
	ERROR_SERVICE_NOT_BIND     = errors.New("service not bind")
)

type server struct{}

func (s *server) init() {
}

// PIPELINE #1 stream receiver
// this function is to make the stream receiving SELECTABLE
func (s *server) recv(stream pb.RoomService_StreamServer, sess_die chan struct{}) chan *pb.Room_Frame {
	ch := make(chan *pb.Room_Frame, 1)
	go func() {
		defer func() {
			close(ch)
		}()
		for {
			in, err := stream.Recv()
			if err == io.EOF { // client closed
				return
			}

			if err != nil {
				log.Error(err)
				return
			}
			select {
			case ch <- in:
			case <-sess_die:
			}
		}
	}()
	return ch
}

// PIPELINE #2 stream processing
// the center of room logic
func (s *server) Stream(stream pb.RoomService_StreamServer) error {
	defer helper.PrintPanicStack()
	// session init
	var sess Session
	sess_die := make(chan struct{})
	ch_agent := s.recv(stream, sess_die)
	ch_ipc := make(chan *pb.Room_Frame, DEFAULT_CH_IPC_SIZE)

	defer func() {
		registry.Unregister(sess.UserId, ch_ipc)
		close(sess_die)
		log.Debug("stream end:", sess.UserId)
	}()

	// read metadata from context
	md, ok := metadata.FromContext(stream.Context())
	if !ok {
		log.Error("cannot read metadata from context")
		return ERROR_INCORRECT_FRAME_TYPE
	}
	// read key
	if len(md["userid"]) == 0 {
		log.Error("cannot read key:userid from metadata")
		return ERROR_INCORRECT_FRAME_TYPE
	}
	// parse userid
	userid, err := strconv.Atoi(md["userid"][0])
	if err != nil {
		log.Error(err)
		return ERROR_INCORRECT_FRAME_TYPE
	}

	// register user
	sess.UserId = int32(userid)
	registry.Register(sess.UserId, ch_ipc)
	log.Debug("userid", sess.UserId, "logged in")

	// >> main message loop <<
	for {
		select {
		case frame, ok := <-ch_agent: // frames from agent
			if !ok { // EOF
				return nil
			}
			switch frame.Type {
			case pb.Room_Message: // the passthrough message from client->agent->room
				// locate handler by proto number
				c := int16(binary.BigEndian.Uint16(frame.Message[:2]))
				handle := client_handler.Handlers[c]
				if handle == nil {
					log.Error("service not bind:", c)
					return ERROR_SERVICE_NOT_BIND

				}

				// handle request
				ret := handle(&sess, frame.Message[2:])

				// construct frame & return message from logic
				if ret != nil {
					if err := stream.Send(&pb.Room_Frame{Type: pb.Room_Message, Message: ret}); err != nil {
						log.Error(err)
						return err
					}
				}

				// session control by logic
				if sess.Flag&SESS_KICKED_OUT != 0 { // logic kick out
					if err := stream.Send(&pb.Room_Frame{Type: pb.Room_Kick}); err != nil {
						log.Error(err)
						return err
					}
					return nil
				}
			case pb.Room_Ping:
				if err := stream.Send(&pb.Room_Frame{Type: pb.Room_Ping, Message: frame.Message}); err != nil {
					log.Error(err)
					return err
				}
				log.Debug("pong")
			default:
				log.Error("incorrect frame type:", frame.Type)
				return ERROR_INCORRECT_FRAME_TYPE
			}
		case frame := <-ch_ipc: // forward async messages from interprocess(goroutines) communication
			if err := stream.Send(frame); err != nil {
				log.Error(err)
				return err
			}
		}
	}
}
