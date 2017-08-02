package auth

import (
	"github.com/Greckas/SSU-Golang-252-Chat/loger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/Greckas/SSU-Golang-252-Chat/database"
)

//type UserRegistry struct {
//	UserName string `json:"user_name"`
//	Password string `json:"password"`
//	NickName string `json:"nick_name"`
//}

type User struct {
	UserName string `gorm:"primary_key"`
	Password string
	NickName string
}



func NewUser(user *User, UserName, Password string) error {
	db, err := database.GetStorage()
	if err != nil {
		loger.Log.Errorf("Failed to open db", err)
		return err
	}
	if err := db.Where("username = ?", UserName).First(&user).Error; err != nil {
		loger.Log.Errorf("no such user in db", err)
		return err
	} else {
		db.NewRecord(user)
		err := db.Create(&user).Error
		if err != nil {
			loger.Log.Errorf("failed to create new user", err)
			return err
		}
		Login(UserName, Password)

	}
	return nil
}
