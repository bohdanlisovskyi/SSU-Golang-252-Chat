package auth

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/Greckas/SSU-Golang-252-Chat/database"
	"github.com/Greckas/SSU-Golang-252-Chat/loger"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func ByName(UserName string) (*User, error) {
	db, err := database.GetStorage()
	if err != nil {
		loger.Log.Errorf("Failed to open db", err)
		return nil, err
	}
	res, err := byQuery(db.Where("username=?", UserName))
	if err != nil {
		loger.Log.Errorf("Search in db failed")
		return nil, err
	}
	return res, err
}

func Login(username, password string) error {
	foundUser, err := ByName(username)
	if err != nil {
		loger.Log.Errorf("No user with that Username")
		return err
	}

	err2 := bcrypt.CompareHashAndPassword(
		[]byte(foundUser.Password),
		[]byte(password))
	if err2 != nil {
		loger.Log.Errorf("Invalid password", err2)
		return err2
	}

	return nil
}

func RandToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func byQuery(db *gorm.DB) (*User, error) {
	ret := &User{}
	err := db.First(ret).Error
	if err != nil {
		loger.Log.Errorf("Failed to find user in DB")
		return nil, err
	}
	return ret, err
}
