package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username      string       `gorm:"unique;not null"`
	WalletAddress string       `gorm:"unique;not null"`
	Courses       []Course      `gorm:"foreignkey:courseID"`
	LearningPath  []LearningPath `gorm:"many2many:users_learningpath"`
	DisplayName   string
}

type UserRepo interface {
	Create(u *User) error
	Update(u *User) error
	UpdateLearningPath(u *User) error
	GetByID(id uint) (*User, error)
	GetByUsername(username string) (*User, error)
	GetByWalletAddr(addr string) (*User, error)
}
