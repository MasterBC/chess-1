package main

import (
"chess/common/define"
"chess/common/log"
"errors"
"golang.org/x/net/context"
"regexp"
. "chess/srv/srv-task/proto"
"chess/srv/srv-task/redis"
    "encoding/json"
)

var (
    ERROR_METHOD_NOT_SUPPORTED = errors.New("method not supoorted")
)
var (
    uuid_regexp = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
)

type server struct {
}

func (s *server) init() {
}


func (s *server) GameOver(ctx context.Context,args *GameInfoArgs) (*BaseRes,error){
    log.Debug("gameOver receive.")
     //判断数据是否是否收到
    if args.Winner != 0 {
	//存入redis 队列
	dataByte,err:=json.Marshal(args)
	if err != nil {
	    log.Errorf(" GameOver err %s",err)
	    return  &BaseRes{Ret:0,Msg:"recive fail."},err
	}
	data:=string(dataByte)
         task_redis.Redis.Task.Lpush(define.TaskLoopHandleGameOverRedisKey,data)
	return &BaseRes{Ret:1,Msg:""},nil
	
    }
    return  &BaseRes{Ret:0,Msg:"recive fail."},nil
}

func (s *server) PlayerEvent(ctx context.Context,args *PlayerActionArgs)(*BaseRes,error){
    log.Debug("PlayerEvent receive.")
    //判断数据是否是否收到
    if args.Id != 0 {
	//存入redis 队列
	dataByte,err:=json.Marshal(args)
	if err != nil {
	    log.Errorf("PlayerEvent err %s",err)
	    return  &BaseRes{Ret:0,Msg:"recive fail."},err
	}
	data:=string(dataByte)
	task_redis.Redis.Task.Lpush(define.TaskLoopHandlePlayerEventRedisKey,data)
	return &BaseRes{Ret:1,Msg:""},nil
    }
    return  &BaseRes{Ret:0,Msg:"recive fail."},nil
}


