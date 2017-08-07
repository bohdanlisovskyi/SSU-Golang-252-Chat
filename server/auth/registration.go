package auth

import (
	"regexp"

	"errors"

	"github.com/8tomat8/SSU-Golang-252-Chat/database"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func RegisterNewUser(user *messageService.User) (*messageService.User, string, error) {
	if checkUsername(user.UserName) != true {
		err := errors.New("input username not valid")
		return nil, "", err
	}
	if checkPassword(user.Password) != true {
		err := errors.New("input password not valid")
		return nil, "", err
	}
	db, err := database.GetStorage()
	if err != nil {
		loger.Log.Errorf("Failed to open db", err)
		return nil, "", err
	}
	err = db.Create(&user).Error
	if err != nil {
		loger.Log.Errorf("failed to create new user or user exist", err)
		return nil, "", err
	}
	user, tok, err := Login(user.UserName, user.Password)
	if err != nil {
		loger.Log.Errorf("Login failed", err)
		return nil, "", err
	}
	return user, tok, nil
}

func checkPassword(password string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", password); !ok {
		return false
	}
	return true
}

func checkUsername(username string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", username); !ok {
		return false
	}
	return true
}
