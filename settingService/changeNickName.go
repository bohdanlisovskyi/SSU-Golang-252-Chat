package settingService

import (
	"encoding/json"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"errors"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	//import 	"github.com/8tomat8/SSU-Golang-252-Chat/database"

)

// ChangeNickNameRequestBody is a custom body for ChangeNickNameRequest
type ChangeNickNameRequestBody struct {
	NickName string `json:"nick_name"`
}

// UnmarshalNickNameRequest function unmarshals request for changing NickName into ChangeNickNameRequestBody struct
// and retrieves value of NickName.
// Function returns: if succeed - NickName value to be stored in users table, nil,
// if failed - empty string, err
func UnmarshalChangeNickNameRequestBody(request messageService.Message) (string, error) {
	var body *ChangeNickNameRequestBody
	err := json.Unmarshal(request.Body, &body)
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return "", err
	}
	nickName := body.NickName
	if nickName == "" {
		err := errors.New("NickName value is empty")
		loger.Log.Errorf("Error has occurred: ", err)
		return "", err
	}
	return nickName, nil
}

// ChangeNickName perform changing users NickName value in users table.
// Function returns: if succeed - true and nil, if failed - false and error
func ChangeNickName(request messageService.Message) (bool, error) {
	userName := request.Header.UserName
	if userName == "" {
		err := errors.New("User name value is empty")
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	nickName, err := UnmarshalChangeNickNameRequestBody(request)
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
	// UPDATE users SET nick_name = "nickName value from request body"
	// WHERE user_name = "userName value from request header"
	db.Model(&User).Where("user_name = ?", userName).Update("nick_name", nickName)
	if db.Error != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	return true, nil
}
