package auth

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/8tomat8/SSU-Golang-252-Chat/database"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	"golang.org/x/crypto/bcrypt"
	"github.com/gorilla/websocket"
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

	// Generate "hash" to check from user password
	hash, err := bcrypt.GenerateFromPassword([]byte(foundUser.Password), bcrypt.DefaultCost)
	if err != nil {
		loger.Log.Errorf("hash generation failed!", err)
		return nil, "", err
	}

	userPassword2 := password
	hashFromDatabase := hash

	// Comparing the password with the hash
	if err := bcrypt.CompareHashAndPassword(hashFromDatabase, []byte(userPassword2)); err != nil {
		loger.Log.Errorf("Password check failed!", err)
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
