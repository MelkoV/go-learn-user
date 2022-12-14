package api

import (
	"context"
	"github.com/MelkoV/go-learn-common/model"
	pb "github.com/MelkoV/go-learn-proto/proto/user"
)

func (s *Api) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	l := s.l.AddSubCategory("login")
	l.WithUuid(in.Uuid)
	l.Info("incoming request: %v", in)

	var user model.User
	s.GetDb().First(&user)

	l.Debug("user is %v", user)

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
