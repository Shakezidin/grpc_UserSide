package service

import (
	"github.com/shakezidin/pkg/DOM"
	adminpb "github.com/shakezidin/pkg/adminpb/pb"
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

func (u *UserService) FetchAllSUserSvc() ([]*DOM.User, error) {
	users, err := u.repo.FetchAllSUserRepo()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserService) DeleteUserSvc(id uint64) (string, error) {
	user, err := u.repo.FindUserbyId(uint(id))
	if err != nil {
		return "usernot found", err
	}

	err = u.repo.DeleteUserById(user.ID)
	if err != nil {
		return "usernot found", err
	}

	return user.UserName, nil
}

func (u *UserService) SearchUserSvc(username string) ([]*DOM.User, error) {
	users, err := u.repo.FindAllUserByUsername(username)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserService) EditUserSvc(p *adminpb.Users) (*adminpb.UserResponse, error) {
	user, err := u.repo.FindUserbyId(uint(p.Id))
	if err != nil {
		return nil, err
	}
	err = u.repo.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return &adminpb.UserResponse{
		Status:   "Success",
		Username: user.UserName,
	}, nil
}

func NewUserService(repos inter.UserRepositoryinter) interr.UserServiceInter {
	return &UserService{
		repo: repos,
	}
}
