package models

import (
	//	"github.com/authenter/nyx/pkg/errors"
	//	"github.com/authenter/nyx/pkg/logger"
	//	"gorm.io/gorm/logger"

	//	"github.com/google/uuid"
	"gorm.io/gorm"

	"time"
)

type User struct {
	ID        string    `json:"id" gorm:"primarykey"`
	Email     string    `json:"email" gorm:"unique"`
	Hash      string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	//	Groups    []*Group  `json:"-" gorm:"many2many:groups_users;"`
}

type UserBeforeSave struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) GetAll(db *gorm.DB) (*[]User, error) {
	var err error
	var users []User

	err = db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return &users, nil
}
