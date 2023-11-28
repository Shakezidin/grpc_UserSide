package server

import (
	"fmt"
	"log"
	"net"

	"github.com/shakezidin/config"
	"github.com/shakezidin/pkg/handler"
	pb "github.com/shakezidin/pkg/pb"
	"google.golang.org/grpc"
)

func NewGrpcServer(cfg *config.Config, handlr *handler.UserHandler) error {
	log.Println("connecting to gRPC server")
	addr := fmt.Sprintf(":%s", cfg.GRPCUSERPORT)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println("error Connecting to gRPC server")
		return err
	}
	grp := grpc.NewServer()
	pb.RegisterUserServiceServer(grp,handlr)
	if err != nil {
		log.Println("error connecting to gRPC server")
		return err
	}

	log.Printf("listening on gRPC server %v", cfg.GRPCUSERPORT)
	err = grp.Serve(lis)
	if err != nil {
		log.Println("error connecting to gRPC server")
		return err
	}
	return nil
}
