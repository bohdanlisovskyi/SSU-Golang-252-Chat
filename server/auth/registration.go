package auth

import (
	"github.com/8tomat8/SSU-Golang-252-Chat/database"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func RegisterNewUser(user *messageService.User) (*messageService.User, error) {
	db, err := database.GetStorage()
	if err != nil {
		loger.Log.Errorf("Failed to open db", err)
		return nil, err
	}
	err = db.Create(&user).Error
	if err != nil {
		loger.Log.Errorf("failed to create new user or user exist", err)
		return nil, err
	}
	_,_, err = Login(user.UserName, user.Password)
	if err != nil{
		loger.Log.Errorf("Login failed", err)
		return nil, err
	}

	return user, nil
}
