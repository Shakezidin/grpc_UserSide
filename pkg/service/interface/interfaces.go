package interfaces

import (
	"github.com/shakezidin/pkg/DOM"
	pb "github.com/shakezidin/pkg/pb"
	adminpb "github.com/shakezidin/pkg/adminpb/pb"
)

type UserServiceInter interface {
	SignupService(pbuser *pb.SignupRequest) (*DOM.User, error)
	LoginService(pbuser *pb.LoginRequest) (*DOM.User, error)
	FetchAllSUserSvc() ([]*DOM.User, error)
	DeleteUserSvc(id uint64) (string, error)
	SearchUserSvc(username string) ([]*DOM.User, error)
	EditUserSvc(p *adminpb.Users) (*adminpb.UserResponse, error)
}
