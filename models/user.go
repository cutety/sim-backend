package models

import (
	"github.com/jinzhu/gorm"
	"sim-backend/extension"
)

var MUser User

type User struct {
	gorm.Model
	UserID string `gorm:"column:user_id;type:varchar(20)" json:"user_id"`
	UserName string `gorm:"column:username;type:varchar(20)" json:"username"`
	Password string `gorm:"column:password;type:varchar(20)" json:"password"`
	Role string `gorm:"column:role;type:int;DEFAULT:2" json:"role"`
}

func (*User) TableName() string {
	return "users"
}

func (*User) GetUserByUserID(userID string) (*User, error) {
	user := &User{}
	err := extension.DB.Where("user_id = ?", userID).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return &User{}, nil
	}
	return user, err
}

func (*User) GetUserByID(id int) (*User, error) {
	user := &User{}
	err := extension.DB.Where("id = ?", id).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return &User{}, nil
	}
	return user, err
}

func (*User) UpdatePasswordById(id int, password string) error {
	return extension.DB.Table(MUser.TableName()).Where("id = ?", id).Update("password", password).Error
}

func (*User) Total() (*int, error) {
	var total *int
	err := extension.DB.Table(MUser.TableName()).Count(&total).Error
	return total, err
}