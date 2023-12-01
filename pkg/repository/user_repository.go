package repository

import (
	"github.com/shakezidin/pkg/DOM"
	inter "github.com/shakezidin/pkg/repository/interface"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (u *UserRepository) CreateUserRepo(user *DOM.User) error {
	result := u.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserRepository) FindUserbyUsername(username string) (*DOM.User, error) {
	var user DOM.User
	result := u.db.Where("user_name = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u *UserRepository) FetchAllSUserRepo() ([]*DOM.User, error) {
	var users []*DOM.User
	result := u.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (u *UserRepository) FindUserbyId(id uint) (*DOM.User, error) {
	var user *DOM.User
	if err := u.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) DeleteUserById(id uint) error {
	var user *DOM.User
	if err := u.db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) FindAllUserByUsername(username string) ([]*DOM.User, error) {
	var users []*DOM.User
	if err := u.db.Where("user_name ILIKE ?", "%"+username+"%").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepository) UpdateUser(user *DOM.User) error {
	if err := u.db.Model(&DOM.User{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func NewUserRepository(db *gorm.DB) inter.UserRepositoryinter {
	return &UserRepository{
		db: db,
	}
}
