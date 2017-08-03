package auth

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/Greckas/SSU-Golang-252-Chat/database"
	"github.com/Greckas/SSU-Golang-252-Chat/loger"
	"github.com/Greckas/SSU-Golang-252-Chat/messageService"
	"golang.org/x/crypto/bcrypt"
)

//+token
func getUserbyName(UserName string) (*messageService.User, error) {

	db, err := database.GetStorage()
	if err != nil {
		loger.Log.Errorf("Failed to open db", err)
		return nil, err
	}
	db_search := db.Where("username=?", UserName)
	if db_search == nil {
		loger.Log.Errorf("Search in db failed")
		return nil, err
	}
	ret := &messageService.User{}
	err = db_search.First(ret).Error
	if err != nil {
		loger.Log.Errorf("Failed to find user in DB")
		return nil, err
	}

	return ret, err
}

func Login(username, password string) (*messageService.User, string, error) {
	foundUser, err := getUserbyName(username)
	if err != nil {
		loger.Log.Errorf("No user with that Username")
		return nil, "", err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(foundUser.Password),
		[]byte(password))
	if err != nil {
		loger.Log.Errorf("Invalid password", err)
		return nil, "", err
	}
	token := randToken()
	return foundUser, token, nil
}

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
