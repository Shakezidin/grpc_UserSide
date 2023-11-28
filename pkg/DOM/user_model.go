package DOM

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"username" gorm:"unique;not null"`
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"unique;not null"`
}

type Login struct {
	Username string
	Password string
}
