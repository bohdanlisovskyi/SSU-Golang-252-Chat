package auth

import (
	"github.com/8tomat8/SSU-Golang-252-Chat/database"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func RegisterNewUser(user *messageService.User) (*messageService.User, string, error) {
	db, err := database.GetStorage()
	if err != nil {
		loger.Log.Errorf("Failed to open db", err)
		return nil, "", err
	}
	if err := db.Where("username = ?", user.UserName).Error; err != nil {
		loger.Log.Errorf("no such user in db", err)
		return nil, "", err
	}

	err = db.Create(&user).Error
	if err != nil {
		loger.Log.Errorf("failed to create new user", err)
		return nil, "", err
	}
	token := randToken()
	Login(user.UserName, user.Password)

	return user, token, nil
}
