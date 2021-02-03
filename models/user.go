package models

import (
	"github.com/jinzhu/gorm"
	"sim-backend/extension"
	"time"
)

var MUser User

type User struct {
	ID        uint `gorm:"primary_key"`
	UserID string `gorm:"column:user_id;type:varchar(20)" json:"user_id"`
	Username string `gorm:"column:username;type:varchar(20)" json:"username"`
	Password string `gorm:"column:password;type:varchar(20)" json:"password"`
	Role int `gorm:"column:role;type:int;DEFAULT:2" json:"role"`
	CreatedAt time.Time `gorm:"type:timestamp"`
	UpdatedAt time.Time `gorm:"type:timestamp"`
	DeletedAt *time.Time `sql:"index" gorm:"type:timestamp"`
}

func (*User) TableName() string {
	return "users"
}

func (*User) GetUserByUserID(userID string) (*User, error) {
	user := &User{}
	err := extension.DB.Where("user_id = ?", userID).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return &User{}, err
	}
	if err != nil {
		return nil, err
	}
	return user, err
}

func (*User) GetUserByID(id uint) (*User, error) {
	user := &User{}
	err := extension.DB.Where("id = ?", id).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return &User{}, nil
	}
	return user, err
}

func (*User) UpdatePasswordById(id uint, password string) error {
	return extension.DB.Table(MUser.TableName()).Where("id = ?", id).Update("password", password).Error
}

func (*User) Total() (*int, error) {
	var total *int
	err := extension.DB.Table(MUser.TableName()).Count(&total).Error
	return total, err
}