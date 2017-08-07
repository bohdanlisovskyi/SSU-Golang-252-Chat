package settingService

import (
	"encoding/json"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"errors"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	//import 	"github.com/8tomat8/SSU-Golang-252-Chat/database"

)

// ChangePassRequestBody is a custom body for ChangePassRequest
type ChangePassRequestBody struct {
	NewPass string `json:"new_pass"`
}

// UnmarshalChangePassRequestBody function unmarshals request for changing Password into ChangePassRequestBody struct
// and retrieves value of NewPass.
// Function returns: if succeed - NewPass value to be stored in users table, nil,
// if failed - empty string, err
func UnmarshalChangePassRequestBody(request messageService.Message) (string, error) {
	var body *ChangePassRequestBody
	err := json.Unmarshal(request.Body, &body)
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return "", err
	}
	newPass := body.NewPass
	if newPass == "" {
		err := errors.New("Password value is empty")
		loger.Log.Errorf("Error has occurred: ", err)
		return "", err
	}
	return newPass, nil
}

// UserResult struct is a copy of Users struct from database package,
// it is used for storing result of querying from users table
type UserResult struct {
	UserName string
	NickName string
	Password string
	Birthday int
	AboutUser string
}

// ChangePass perform changing users Password value in users table.
// Function returns: if succeed - true and nil, if failed - false and error
func ChangePass(request messageService.Message) (bool, error) {
	userName := request.Header.UserName
	if userName == "" {
		err := errors.New("User name value is empty")
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	newPass, err := UnmarshalChangePassRequestBody(request)
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	db, err := GetStorage() // common gorm-connection from database package
	defer db.Close()
	if err != nil {
		loger.Log.Errorf("DB error has occurred: ", err)
		return false, err
	}
	var result UserResult //variable for storing result of querying into UserResult struct
	// SELECT user_name, nick_name, password, birthday, about_user FROM users
	// WHERE main_user = "mainUser value from request body"
	// AND contact_user = "contactUser value from request body"
	db.Table("users").
		Select("user_name, nick_name, password, birthday, about_user").
		Where("user_name = ?", userName).Scan(&result)
	oldPass := result.Password
	if newPass == oldPass {
		loger.Log.Warnf(" Password is the same as old")
		return false, nil
	}
	// UPDATE users SET password = "newPass value from request body"
	// WHERE user_name = "userName value from request header"
	db.Model(&User).Where("user_name = ?", userName).Update("password", newPass)
	if db.Error != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	return true, nil
}
