package handler

import (
	"context"

	adminpb "github.com/shakezidin/pkg/adminpb/pb"
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

func (h *AdminHandler)	DeleteUser(ctx context.Context,p *adminpb.DeleteUserById) (*adminpb.AdminResult, error){
	
}
// 	CreateUser(context.Context, *UserCreate) (*Result, error)
// 	SearchUser(context.Context, *UserRequest) (*SearchResponse, error)
// 	EditUser(context.Context, *Users) (*UserResponse, error)
