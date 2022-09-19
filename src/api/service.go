package api

import (
	"fmt"
	"github.com/MelkoV/go-learn-logger/logger"
	pb "github.com/MelkoV/go-learn-proto/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Api struct {
	pb.UnimplementedUserServiceServer
	l *logger.CategoryLogger // @TODO change to interface
}

func NewApi(l *logger.CategoryLogger) *Api {
	return &Api{
		l: l,
	}
}

func Serve(port int, l *logger.CategoryLogger) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		l.Format("init", "00000000-0000-0000-0000-000000000000", "failed listen port %d: %s", port, err.Error()).Fatal()
	}

	l.Format("init", "00000000-0000-0000-0000-000000000000", "running API server on port %d", port).Info()

	server := grpc.NewServer()
	api := NewApi(l)
	pb.RegisterUserServiceServer(server, api)

	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		l.Format("init", "00000000-0000-0000-0000-000000000000", "failed to serve: %s", err.Error()).Fatal()
	}
}
