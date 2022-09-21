package api

import (
	"context"
	pb "github.com/MelkoV/go-learn-proto/proto/user"
)

func (s *Api) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	s.l.Format("login", in.Uuid, "incoming request: %v", in).Info()
	cookies := []*pb.Cookie{
		&pb.Cookie{
			Name:   "test1",
			Value:  "value1",
			MaxAge: 86400,
		},
		&pb.Cookie{
			Name:   "test2",
			Value:  "value2",
			MaxAge: 0,
		},
	}
	return &pb.LoginResponse{
		User: &pb.User{
			Id:   100,
			Name: "Test",
		},
		Cookie: cookies,
	}, nil
}
