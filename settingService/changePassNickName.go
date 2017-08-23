package settingService

import (
	"encoding/json"
	"errors"

	"github.com/8tomat8/SSU-Golang-252-Chat/database"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
)

// ChangeUserDataRequestBody is a custom body for request to change nick_name or password
type ChangeUserDataRequestBody struct {
	Password string `json:"password"`
	NickName string `json:"nick_name"`
}

// UnmarshalUserDataRequestBody function retrieves values of password and nickName from request body
func UnmarshalUserDataRequestBody(request *messageService.Message) (pass, nickName string, err error) {
	var body *ChangeUserDataRequestBody
	err = json.Unmarshal(request.Body, &body)
	if err != nil {
		loger.Log.Error(err)
		return
	}
	if body.NickName == "" {
		err := errors.New("NickName value is empty")
		loger.Log.Error(err)
		return
	}
	if body.Password == "" {
		err := errors.New("Password value is empty")
		loger.Log.Error(err)
		return
	}
	return body.Password, body.NickName, nil
}

// GetUserName retrieves value of userName from request header
func GetUserName(request *messageService.Message) (userName string, err error) {
	if request.Header.UserName == "" {
		err := errors.New("User name value is empty")
		loger.Log.Error(err)
		return
	}
	return request.Header.UserName, nil
}

// ChangeNickName perform changing users NickName value in authentifications table.
func ChangeNickName(request *messageService.Message) (ok bool, err error) {
	userName, err := GetUserName(request)
	if err != nil {
		return
	}
	pass, nickName, err := UnmarshalUserDataRequestBody(request)
	if err != nil {
		return
	}
	//authentifications - table's name in DB with required fields
	authentification := messageService.Authentification{UserName: userName, NickName: nickName, Password: pass}
	db, err := database.GetStorage()
	if err != nil {
		return
	}
	db.Model(&authentification).Update("nick_name", nickName)
	if db.Error != nil {
		loger.Log.Error(err)
		return
	}
	return true, nil
}

// ChangePass perform changing users Password value in Authentifications table.
func ChangePass(request *messageService.Message) (ok bool, err error) {
	userName, err := GetUserName(request)
	if err != nil {
		return
	}
	pass, nickName, err := UnmarshalUserDataRequestBody(request)
	if err != nil {
		return
	}
	db, err := database.GetStorage()
	if err != nil {
		return
	}
	//authentification - table's name in DB with required fields
	authentification := messageService.Authentification{UserName: userName, NickName: nickName, Password: pass}
	db.Model(&authentification).Update("password", pass)
	if db.Error != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return
	}
	return true, nil
}
