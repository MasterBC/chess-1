package main

import (
	"errors"
	"fmt"
	"snowflake/etcdclient"
	pb "snowflake/proto"
	"strconv"
	"sync"
	"time"

	cli "gopkg.in/urfave/cli.v2"

	log "github.com/Sirupsen/logrus"
	etcd "github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

const (
	BACKOFF    = 100  // max backoff delay millisecond
	CONCURRENT = 128  // max concurrent connections to etcd
	UUID_QUEUE = 1024 // uuid process queue
)

const (
	TS_MASK         = 0x1FFFFFFFFFF // 41bit
	SN_MASK         = 0xFFF         // 12bit
	MACHINE_ID_MASK = 0x3FF         // 10bit
)

type server struct {
	pkroot     string
	uuidkey    string
	machine_id uint64 // 10-bit machine id
	ch_proc    chan chan uint64
	muNext     sync.Mutex
}

func (s *server) init(c *cli.Context) {
	etcdclient.Init(c)
	s.ch_proc = make(chan chan uint64, UUID_QUEUE)
	// shifted machine id
	s.machine_id = (uint64(c.Int("machine-id")) & MACHINE_ID_MASK) << 12
	s.pkroot = c.String("pk-root")
	s.uuidkey = c.String("uuid-key")
	go s.uuid_task()
}

// get next value of a key, like auto-increment in mysql
func (s *server) Next(ctx context.Context, in *pb.Snowflake_Key) (*pb.Snowflake_Value, error) {
	s.muNext.Lock()
	defer s.muNext.Unlock()
	client := etcdclient.KeysAPI()
	key := s.pkroot + "/" + in.Name
	for {
		// get the key
		resp, err := client.Get(context.Background(), key, nil)
		if err != nil {
			log.Error(err)
			return nil, errors.New("Key not exists, need to create first")
		}

		// get prevValue & prevIndex
		prevValue, err := strconv.Atoi(resp.Node.Value)
		if err != nil {
			log.Error(err)
			return nil, errors.New("marlformed value")
		}
		prevIndex := resp.Node.ModifiedIndex

		// CompareAndSwap
		resp, err = client.Set(context.Background(), key, fmt.Sprint(prevValue+1), &etcd.SetOptions{PrevIndex: prevIndex})
		if err != nil {
			log.Warn(err)
			continue
		}
		return &pb.Snowflake_Value{int64(prevValue + 1)}, nil
	}
}

// generate an unique uuid
func (s *server) GetUUID(context.Context, *pb.Snowflake_NullRequest) (*pb.Snowflake_UUID, error) {
	req := make(chan uint64, 1)
	s.ch_proc <- req
	return &pb.Snowflake_UUID{<-req}, nil
}

// uuid generator
func (s *server) uuid_task() {
	var sn uint64     // 12-bit serial no
	var last_ts int64 // last timestamp
	for {
		ret := <-s.ch_proc
		// get a correct serial number
		t := ts()
		if t < last_ts { // clock shift backward
			log.Warn("clock shift happened, waiting until the clock moving to the next millisecond.")
			t = s.wait_ms(last_ts)
		}

		if last_ts == t { // same millisecond
			sn = (sn + 1) & SN_MASK
			if sn == 0 { // serial number overflows, wait until next ms
				t = s.wait_ms(last_ts)
			}
		} else { // new millsecond, reset serial number to 0
			sn = 0
		}
		// remember last timestamp
		last_ts = t

		// generate uuid, format:
		//
		// 0		0.................0		0..............0	0........0
		// 1-bit	41bit timestamp			10bit machine-id	12bit sn
		var uuid uint64
		uuid |= (uint64(t) & TS_MASK) << 22
		uuid |= s.machine_id
		uuid |= sn
		ret <- uuid
	}
}

// wait_ms will wait untill last_ts
func (s *server) wait_ms(last_ts int64) int64 {
	t := ts()
	for t < last_ts {
		time.Sleep(time.Duration(last_ts-t) * time.Millisecond)
		t = ts()
	}
	return t
}

// get timestamp
func ts() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
