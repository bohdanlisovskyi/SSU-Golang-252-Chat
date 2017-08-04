package auth

import (
	"crypto/rand"
	"encoding/base64"

	"fmt"

	"github.com/8tomat8/SSU-Golang-252-Chat/database"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	"golang.org/x/crypto/bcrypt"
)

//+token
func getUserByName(UserName string) (*messageService.User, error) {

	db, err := database.GetStorage()
	if err != nil {
		loger.Log.Errorf("Failed to open db", err)
		return nil, err
	}
	db_search := db.Where("user_name=?", UserName)
	ret := &messageService.User{}
	err = db_search.First(ret).Error
	if err != nil {
		loger.Log.Errorf("Failed to find user in DB")
	}

	return ret, err
}

func Login(username, password string) (*messageService.User, string, error) {
	foundUser, err := getUserByName(username)
	if err != nil {
		loger.Log.Errorf("No user with that Username")
		return nil, "", err
	}
	db_p := []byte(foundUser.Password)
	cl_p := []byte(password)
	fmt.Println(db_p, cl_p)
	err = bcrypt.CompareHashAndPassword(
		db_p,
		cl_p)
	if err != nil {
		loger.Log.Errorf("Invalid password", err)
		return nil, "", err
	}
	token := randToken()
	//write message with header & body
	return foundUser, token, nil
}

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
