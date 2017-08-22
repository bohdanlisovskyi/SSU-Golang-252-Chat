package settingService

import (
	"encoding/json"
	"errors"

	"github.com/8tomat8/SSU-Golang-252-Chat/database"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
)

// ChangeNickNamePassRequestBody is a custom body for request to change nick_name or password
type ChangeNickNamePassRequestBody struct {
	Password string `json:"password"`
	NickName string `json:"nick_name"`
}

// UnmarshalNickNameRequest function retrieves values of password and nickName from request body
func UnmarshalRequestBody(request *messageService.Message) (pass string, nickName string, err error) {
	var body *ChangeNickNamePassRequestBody
	err = json.Unmarshal(request.Body, &body)
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return "", "", err
	}
	pass = body.Password
	nickName = body.NickName
	if nickName == "" {
		err := errors.New("NickName value is empty")
		loger.Log.Errorf("Error has occurred: ", err)
		return "", "", err
	}
	if pass == "" {
		err := errors.New("Password value is empty")
		loger.Log.Errorf("Error has occurred: ", err)
		return "", "", err
	}
	return pass, nickName, nil
}

// GetUserName retrieves value of userName from request header
func GetUserName(request *messageService.Message) (string, error) {
	userName := request.Header.UserName
	if userName == "" {
		err := errors.New("User name value is empty")
		loger.Log.Errorf("Error has occurred: ", err)
		return "", err
	}
	return userName, nil
}

// ChangeNickName perform changing users NickName value in authentifications table.
func ChangeNickName(request *messageService.Message) (bool, error) {
	userName, err := GetUserName(request)
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	pass, nickName, err := UnmarshalRequestBody(request)
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	//authentifications - table's name in DB with needed fields
	authentification := messageService.Authentification{UserName: userName, NickName: nickName, Password: pass}
	db, err := database.GetStorage()
	if err != nil {
		loger.Log.Errorf("DB error has occurred: ", err)
		return false, err
	}
	db.Model(&authentification).Update("nick_name", nickName)
	if db.Error != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	return true, nil
}

//IsPassTheSame checks old and new passwords
func IsPassTheSame(request *messageService.Message) (bool, error) {
	db, err := database.GetStorage()
	if err != nil {
		loger.Log.Errorf("DB error has occurred: ", err)
		return true, err
	}
	userName, err := GetUserName(request)
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return true, err
	}
	newPass, _, _ := UnmarshalRequestBody(request)
	var result *messageService.Authentification //variable for storing result of querying by Authentifications table
	// SELECT user_name, nick_name, password, birthday, about_user FROM Authentifications
	// WHERE user_name = "userName value from request header"
	db.Table("authentifications").
		Select("user_name, nick_name, password").
		Where("user_name = ?", userName).
		Scan(&result)
	oldPass := result.Password
	if newPass == oldPass {
		loger.Log.Warnf(" Password is the same as old")
		return true, nil
	}
	return false, nil
}

// ChangePass perform changing users Password value in Authentifications table.
func ChangePass(request *messageService.Message) (bool, error) {
	userName, err := GetUserName(request)
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	pass, nickName, err := UnmarshalRequestBody(request)
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	db, err := database.GetStorage() // common gorm-connection from database package
	if err != nil {
		loger.Log.Errorf("DB error has occurred: ", err)
		return false, err
	}
	isPassSame, err := IsPassTheSame(request)
	if isPassSame {
		loger.Log.Warnf(" Password is the same as old")
		return false, nil
	}
	//authentification - table's name in DB with needed fields
	authentification := messageService.Authentification{UserName: userName, NickName: nickName, Password: pass}
	db.Model(&authentification).Update("password", pass)
	if db.Error != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	return true, nil
}
