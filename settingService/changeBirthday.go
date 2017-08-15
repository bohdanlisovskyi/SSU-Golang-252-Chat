package settingService

import (
	"encoding/json"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"errors"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	"github.com/8tomat8/SSU-Golang-252-Chat/database"
)

// ChangeBirthdayRequestBody is a custom body for ChangeBirthdayRequest
type ChangeBirthdayRequestBody struct {
	Birthday int `json:"birthday"`
}

// UnmarshalBirthdayRequest function unmarshals request for changing birthday into ChangeBirthdayRequestBody struct
// and retrieves value of birthday.
// Function returns: if succeed - birthday value to be stored in users table, nil,
// if failed - -1, err
func UnmarshalChangeBirthdayRequestBody(request *messageService.Message) (int, error) {
	var body *ChangeBirthdayRequestBody
	err := json.Unmarshal(request.Body, &body)
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return -1, err
	}
	birthday := body.Birthday
	return birthday, nil
}

// ChangeBirthday perform changing users birthday value in users table.
// Function returns: if succeed - true and nil, if failed - false and error
func ChangeBirthday(request *messageService.Message) (bool, error) {
	userName := request.Header.UserName
	if userName == "" {
		err := errors.New("User name value is empty")
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	birthday, err := UnmarshalChangeBirthdayRequestBody(request)
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	db, err := database.GetStorage() // common gorm-connection from database package
	if err != nil {
		loger.Log.Errorf("DB error has occurred: ", err)
		return false, err
	}
	// UPDATE users SET birthday = "birthday value from request body"
	// WHERE user_name = "userName value from request header"
	db.Model(&User).Where("user_name = ?", userName).Update("birthday", birthday)
	if db.Error != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	return true, nil
}
