package server

import (
	"fmt"
	"log"
	"net"

	"github.com/shakezidin/config"
	adminpb "github.com/shakezidin/pkg/adminpb/pb"
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
	pb.RegisterUserServiceServer(grp, handlr)
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

func NewAdminGrpcServer(cfg *config.Config, handlr *handler.AdminHandler) error {
	log.Println("connecting to gRPC server")
	addr := fmt.Sprintf(":%s", cfg.GRPCUSERADMINPORT)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println("error Connecting to gRPC server")
		return err
	}
	grp := grpc.NewServer()
	adminpb.RegisterAdminUserServiceServer(grp, handlr)
	if err != nil {
		log.Println("error connecting to gRPC server")
		return err
	}

	log.Printf("listening on gRPC server %v", cfg.GRPCUSERADMINPORT)
	err = grp.Serve(lis)
	if err != nil {
		log.Println("error connecting to gRPC server")
		return err
	}
	return nil
}

func GrpcServerRun(cfg *config.Config, userHandler *handler.UserHandler, adminHandler *handler.AdminHandler) {
	go NewAdminGrpcServer(cfg, adminHandler)
	err := NewGrpcServer(cfg, userHandler)
	if err != nil {
		log.Fatalf("something went wrong: %v", err)

	}
 

 }
 