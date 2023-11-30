package interfaces

import "github.com/shakezidin/pkg/DOM"

type UserRepositoryinter interface {
	CreateUserRepo(user *DOM.User) error
	FindUserbyUsername(username string)(*DOM.User,error)
	FetchAllSUserRepo()([]*DOM.User,error)
	FindUserbyId(id uint64)(*DOM.User,error)
	DeleteUserById(id uint) error
}
