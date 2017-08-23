package grpc

import (
	pb "chess/api/proto"
	"chess/common/config"
	"chess/common/services"
)

//var (
//    AuthClient pb.AuthServiceClient
//)

func GetAuthGrpc() (AuthClient pb.AuthServiceClient) {
	auth := services.GetService(config.C.GrpcServer[0])
	AuthClient = pb.NewAuthServiceClient(auth)
	return AuthClient
}

func GetCentreGrpc() (CentreClient pb.CentreServiceClient) {
    centre := services.GetService(config.C.GrpcServer[1])
    CentreClient = pb.NewCentreServiceClient(centre)
    return CentreClient
}