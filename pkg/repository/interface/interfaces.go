package interfaces

import "github.com/shakezidin/pkg/DOM"

type UserRepositoryinter interface {
	CreateUserRepo(user *DOM.User) error
	FindUserbyUsername(username string)(*DOM.User,error)
	FetchAllSUserRepo()([]*DOM.User,error)
	FindUserbyId(id uint)(*DOM.User,error)
	DeleteUserById(id uint) error
	FindAllUserByUsername(username string) ([]*DOM.User, error) 
	UpdateUser(user *DOM.User) error
}
