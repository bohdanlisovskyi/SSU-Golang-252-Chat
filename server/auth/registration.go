package auth

import (
	"errors"

	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/jinzhu/gorm"
)

// UserRegistry is a structure for Messager
type UserRegistry struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	NickName string `json:"nick_name"`
}

type User struct {
	UserName string `gorm:"not null;unique_index"`
	Password string
	NickName string
}

type UserGorm struct {
	*gorm.DB
}

func (ug *UserGorm) Create(user *User) error {
	newUser := ug.NewRecord(user)
	if newUser != true {
		err := errors.New("User already exist in database!")
		loger.Log.Errorf("user already exist!", err)
		return err
	} else {
		ug.DB.Create(&user)
		checkUser := ug.NewRecord(user)
		if checkUser != false {
			err := errors.New("Failed to create user!")
			return err
		}
		return nil
	}
}
