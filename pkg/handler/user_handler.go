package handler

import (
	"context"
	"fmt"

	pb "github.com/shakezidin/pkg/pb"
	inter "github.com/shakezidin/pkg/service/interface"
)

type UserHandler struct {
	svc inter.UserServiceInter
	pb.UserServiceServer
}

func NewUserHandler(svc inter.UserServiceInter) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (h *UserHandler) UserSignup(ctx context.Context, p *pb.SignupRequest) (*pb.Result, error) {
	result, err := h.svc.SignupService(p)
	if err != nil {
		return &pb.Result{
			Status:  "false",
			Error:   "User signup Error",
			Message: "",
		}, err
	}
	msg := fmt.Sprintf("User created", result)
	return &pb.Result{
		Status:  "true",
		Error:   "User Signup done",
		Message: msg,
	}, nil
}

func (h *UserHandler) UserLogin(ctx context.Context, p *pb.LoginRequest) (*pb.Result, error) {
	result, err := h.svc.LoginService(p)
	if err != nil {
		msg := fmt.Sprintf("error", result)
		return &pb.Result{
			Status:  "False",
			Error:   "Login Error",
			Message: msg,
		}, err
	}

	msg := fmt.Sprintf("error", result)
	return &pb.Result{
		Status:  "true",
		Error:   "No error",
		Message: msg,
	}, nil
}
