package api

import (
	"context"
	pb "github.com/MelkoV/go-learn-proto/proto/user"
)

func (s *Api) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	s.l.Format("login", "0000", "incoming request: %v", in)
	return &pb.LoginResponse{
		User: &pb.User{
			Id:   100,
			Name: "Test",
		},
	}, nil
}
