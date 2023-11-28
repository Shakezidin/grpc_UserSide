package interfaces

import (
	"github.com/shakezidin/pkg/DOM"
	pb "github.com/shakezidin/pkg/pb"
)

type UserServiceInter interface {
	SignupService(pbuser *pb.SignupRequest) (*DOM.User, error)
	LoginService(pbuser *pb.LoginRequest) (*DOM.User, error)
}
