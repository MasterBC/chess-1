package main

import (
	"chess/common/log"
	. "chess/srv/srv-auth/proto"
	"errors"
	"golang.org/x/net/context"
	"regexp"
	"time"
        "chess/common/auth"
        "chess/common/define"
        "chess/models"
        "chess/common/config"
        "strconv"
        "chess/srv/srv-auth/redis"
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

func (s *server) Auth(ctx context.Context, args *AuthArgs) (*AuthRes, error) {
	// TODO check token
    //判断黑名单
    msg,err:=redis.Redis.Login.Get(args.Token)
    if err == nil {
	return &AuthRes{Ret: define.AuthALreadyLogin, Msg: msg}, nil
    }
    loginData,err :=  auth.AuthLoginToken(args.Token,config.CAuth.TokenSecret)
    if err != nil || strconv.Itoa(int(args.UserId)) != loginData {
	return &AuthRes{Ret:0,Msg:""},err
    }
	log.Debugf("AuthArgs(%+v)", *args)
	return &AuthRes{Ret: 1, Msg: "ok"}, nil
}

func (s *server) TokenInfo(ctx context.Context, args *TokenInfoArgs) (*TokenInfoRes, error) {
	log.Debugf("TokenInfoArgs(%+v)", *args)

	// Get the session
	session, err := models.Session.Get(int(args.UserId))
	if err != nil {
	    return  &TokenInfoRes{Ret:0,Msg:"",Expire: time.Now().Unix()}, nil
	}
	return &TokenInfoRes{Ret:1,Msg:"",Expire:session.Token.Expire}, nil
}

func (s *server) RefreshToken(ctx context.Context, args *RefreshTokenArgs) (*RefreshTokenRes, error) {
	log.Debugf("RefreshTokenArgs(%+v)", *args)
        //查出旧的token,加入没名单
    	session, err := models.Session.Get(int(args.UserId))
	if session != nil {
	    now:=time.Now().Unix()
	    redis.Redis.Login.SetexInt(session.Token.Data,define.AuthALreadyLogin,session.Token.Expire-now)
	}

        result,err :=auth.LoginUser(int(args.UserId),args.AppFrom,args.UniqueId)
	if err != nil {
	    return &RefreshTokenRes{Ret: 0,
		Msg: "ok" ,
	    }, nil
	}

	return &RefreshTokenRes{Ret: 1,
	    Msg: "ok" ,
	    UserId:int32(result.UserId),
	    Token:result.Token ,
	    Expire:result.Expire ,
	    RefreshToken:result.RefreshToken}, nil
}

func (s *server) DestroyToken(ctx context.Context, args *DestroyTokenArgs) (*DestroyTokenRes, error) {
	log.Debugf("RefreshTokenArgs(%+v)", *args)

	return &DestroyTokenRes{Ret: 1}, nil
}
