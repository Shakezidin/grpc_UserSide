package handler

import (
	"context"

	adminpb "github.com/shakezidin/pkg/adminpb/pb"
	pb "github.com/shakezidin/pkg/pb"
	inter "github.com/shakezidin/pkg/service/interface"
)

type AdminHandler struct {
	svc inter.UserServiceInter
	adminpb.AdminUserServiceServer
}

func NewAdminHandler(svc inter.UserServiceInter) *AdminHandler {
	return &AdminHandler{
		svc: svc,
	}
}

func (h *AdminHandler) FetchAllSUser(ctx context.Context, p *adminpb.FetchUsers) (*adminpb.LoginResponce, error) {
	result, err := h.svc.FetchAllSUserSvc()
	if err != nil {
		return nil, err
	}
	var users []*adminpb.Users
	for _, i := range result {
		users = append(users, &adminpb.Users{
			Id:       uint64(i.ID),
			Username: i.UserName,
			Name:     i.Name,
			Email:    i.Email,
			Password: i.Password,
		})
	}
	rslt := &adminpb.LoginResponce{
		Status:    "Success",
		Available: users,
	}
	return rslt, nil
}

func (h *AdminHandler) CreateUser(ctx context.Context, p *adminpb.UserCreate) (*adminpb.AdminResult, error) {
	user := &pb.SignupRequest{
		Username: p.Username,
		Name:     p.Name,
		Email:    p.Email,
		Password: p.Password,
	}

	_, err := h.svc.SignupService(user)
	if err != nil {
		return nil, err
	}
	reslt := &adminpb.AdminResult{
		Status:  "Success",
		Error:   "nil",
		Message: "user created success",
	}
	return reslt, nil
}

func (h *AdminHandler) DeleteUser(ctx context.Context, p *adminpb.DeleteUserById) (*adminpb.AdminResult, error) {
	username, err := h.svc.DeleteUserSvc(p.Id)
	if err != nil {
		return nil, err
	}
	rslt := adminpb.AdminResult{
		Status:  "success",
		Error:   "",
		Message: username,
	}
	return &rslt, nil
}

func (h *AdminHandler) SearchUser(ctx context.Context, p *adminpb.UserRequest) (*adminpb.SearchResponse, error) {
	userss, err := h.svc.SearchUserSvc(p.Username)
	if err != nil {
		return nil, err
	}
	var users []*adminpb.Users
	for _, i := range userss {
		users = append(users, &adminpb.Users{
			Id:       uint64(i.ID),
			Username: i.UserName,
			Name:     i.Name,
			Email:    i.Email,
			Password: i.Password,
		})
	}
	rslt := &adminpb.SearchResponse{
		Status:    "Success",
		Available: users,
	}
	return rslt, nil
}

func (h *AdminHandler) EditUser(ctx context.Context, p *adminpb.Users) (*adminpb.UserResponse, error) {
	result, err := h.svc.EditUserSvc(p)
	if err != nil {
		return nil, err
	}
	return result, nil
}
