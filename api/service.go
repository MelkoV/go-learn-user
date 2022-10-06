package api

import (
	"fmt"
	"github.com/MelkoV/go-learn-logger/logger"
	pb "github.com/MelkoV/go-learn-proto/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net"
	"os"
	"time"
)

var dbStmt *gorm.DB

type Api struct {
	pb.UnimplementedUserServiceServer
	l logger.CategoryLogger
}

func NewApi(l logger.CategoryLogger) *Api {
	return &Api{
		l: l,
	}
}

func (*Api) GetDb() *gorm.DB {
	return dbStmt
}

func Serve(port int, l logger.CategoryLogger, dsn string) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		l.Fatal("failed listen port %d: %s", port, err.Error())
	}

	l.Info("running API server on port %d", port)

	if err := dbConnect(dsn); err != nil {
		l.Fatal("failed to DB connect: %s", err)
		os.Exit(1)
	}

	server := grpc.NewServer()
	api := NewApi(l)
	pb.RegisterUserServiceServer(server, api)

	reflection.Register(server)
	if err := server.Serve(lis); err != nil {
		l.Fatal("failed to serve: %s", err.Error())
	}
}

func dbConnect(dsn string) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	dbStmt = db
	sqlDb, err := db.DB()
	if err != nil {
		return err
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(20)
	sqlDb.SetConnMaxLifetime(time.Hour)
	return nil
}
