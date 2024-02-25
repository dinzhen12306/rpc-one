package server

import (
	"context"
	pb "github.com/dinzhen12306/rpc-one/userrpc"
	"google.golang.org/grpc"
	"net"
)

type UserRpcServer struct {
	pb.UnimplementedStreamGreeterServer
}

func (s *UserRpcServer) SayHello(ctx context.Context, in *pb.SayHelloReq) (*pb.SayHelloResp, error) {

	return &pb.SayHelloResp{
		Greet: "hello" + in.Name,
	}, nil
}

func NewUserRpcServer(port string) {
	grpcServer := grpc.NewServer()
	pb.RegisterStreamGreeterServer(grpcServer, new(UserRpcServer))
	listen, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	reflection.Register(grpcServer)
	err = grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}
}
