package service

import (
	"github.com/shakezidin/pkg/DOM"
	pb "github.com/shakezidin/pkg/pb"
	inter "github.com/shakezidin/pkg/repository/interface"
	interr "github.com/shakezidin/pkg/service/interface"
)

type UserService struct {
	repo inter.UserRepositoryinter
}

func (u *UserService) SignupService(pbuser *pb.SignupRequest) (*DOM.User, error) {
	user := &DOM.User{
		UserName: pbuser.Username,
		Name:     pbuser.Name,
		Email:    pbuser.Email,
		Password: pbuser.Password,
	}
	err := u.repo.CreateUserRepo(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) LoginService(pbuser *pb.LoginRequest) (*DOM.User, error) {
	user, err := u.repo.FindUserbyUsername(pbuser.Username)
	if err != nil {
		return nil, err
	}
	if user.Password != pbuser.Password {
		return nil, nil
	}
	return user, nil
}

func NewUserService(repos inter.UserRepositoryinter) interr.UserServiceInter {
	return &UserService{
		repo: repos,
	}
}
