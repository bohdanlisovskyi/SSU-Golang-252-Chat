package settingService

import (
	"encoding/json"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"errors"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	"github.com/8tomat8/SSU-Golang-252-Chat/database"
)

// ChangeAboutUserRequestBody is a custom body for ChangeAboutUserRequest
type ChangeAboutUserRequestBody struct {
	AboutUser string `json:"about_user"`
}

// UnmarshalAboutUserRequestBody function unmarshals request for changing field about_user(in table users)
// into ChangeBirthdayRequestBody struct and retrieves value of to be stored in about_user field.
// Function returns: if succeed - about_user value to be stored in users table, nil,
// if failed - "", err
func UnmarshalAboutUserRequestBody(request *messageService.Message) (string, error) {
	var body *ChangeAboutUserRequestBody
	err := json.Unmarshal(request.Body, &body)
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return "", err
	}
	aboutUser := body.AboutUser
	if aboutUser == "" {
		err := errors.New("Info hasn't been filled")
		loger.Log.Errorf("Error has occurred: ", err)
		return "", err
	}
	if len(aboutUser) >= 999{ // column about_user in table users has length restriction - VARCHAR(1000)
		err := errors.New("Too many symbols have been filled")
		loger.Log.Errorf("Error has occurred: ", err)
		return "", err
	}
	return aboutUser, nil
}

// ChangeAboutUser perform changing field about_user in users table.
// Function returns: if succeed - true and nil, if failed - false and error
func ChangeAboutUserInfo(request *messageService.Message) (bool, error) {
	userName := request.Header.UserName
	if userName == "" {
		err := errors.New("User name value is empty")
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	/*aboutUser, err := UnmarshalAboutUserRequestBody(request)
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}*/
	db, err := database.GetStorage() // common gorm-connection from database package
	if err != nil {
		loger.Log.Errorf("DB error has occurred: ", err)
		return false, err
	}
	// UPDATE users SET about_user = "aboutUser value from request body"
	// WHERE user_name = "userName value from request header"
//	db.Model(&User).Where("user_name = ?", userName).Update("about_user", aboutUser)
	if db.Error != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	return true, nil
}
