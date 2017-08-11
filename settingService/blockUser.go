package settingService

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"encoding/json"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"errors"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	"github.com/8tomat8/SSU-Golang-252-Chat/database"
)

// BlockUserRequestBody is a custom body for BlockUserRequest
type BlockUserRequestBody struct {
	MainUser    string `json:"main_user"`
	ContactUser string `json:"contact_user"`
	IsBlocked   int `json:"is_blocked"`
}

// UnmarshalBlockUserRequestBody function unmarshals request for blocking of users into BlockUserRequestBody struct,
// checks and retrieves value of is_blocked variable.
// Function returns: if succeed - *BlockUserRequestBody, is_blocked value, nil,
// if failed - nil, nil, err
func UnmarshalBlockUserRequestBody(request *messageService.Message) (*BlockUserRequestBody, int, error) {
	var body *BlockUserRequestBody
	err := json.Unmarshal(request.Body, &body)
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return nil, nil, err
	}
	IsBlocked := body.IsBlocked // IsBlocked value could be int 0 or 1
	if IsBlocked != 0 && IsBlocked != 1 {
		err := errors.New("IsBlocked value is not valid. IsBlocked = " + string(IsBlocked))
		loger.Log.Errorf("IsBlocked value is not 1 or 0:", err)
		return nil, nil, err
	}
	return body, IsBlocked, nil
}

// BlockUnblockUser perform blocking and unblocking of users by clicking the button on UI.
// Button works like a trigger: first push - block user, second push - unblock user.
// Function returns: if succeed - true and nil, if failed - false and error
func BlockUnblockUser(request *messageService.Message) (bool, error) {
	body, IsBlocked, err := UnmarshalBlockUserRequestBody(request)
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	contactUser := body.ContactUser
	mainUser := body.MainUser
	if contactUser == "" || mainUser == "" {
		err := errors.New("Some field or fields are empty")
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	db, err := database.GetStorage() // common gorm-connection from database package
	if err != nil {
		loger.Log.Errorf("DB error has occurred: ", err)
		return false, err
	}
	// UPDATE contacts SET IsBlocked = "IsBlocked value from request body"
	// WHERE main_user = "mainUser value from request body"
	// AND  contact_user = "contactUser value from request body"
	db.Model(&Contact).Where("main_user = ? AND contact_user = ?", mainUser, contactUser).Update("is_blocked", IsBlocked)
	if db.Error != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	return true, nil
}

// ContactResult struct is a copy of Contact struct from database package,
// it is used for storing result of querying from contacts table
type ContactResult struct {
	MainUser    string
	ContactUser string
	IsBlocked   int
}

// IsUserBlocked checks if user is blocked for chatting in contacts table
// Function returns: if succeed - true or false from contacts table(depend is contact blocked or not) and nil, if failed - nil and error
func IsUserBlocked(request *messageService.Message) (isBlocked bool, err error) {
	isBlocked = true
	contactUser := request.Body.ReceiverName
	mainUser := request.Header.UserName
	db, err := database.GetStorage() // common gorm-connection from database package
	if err != nil {
		loger.Log.Errorf("DB error has occurred: ", err)
		return
	}
	var result ContactResult //variable for storing result of querying into ContactResult struct
	// SELECT main_user, contact_user, is_blocked FROM contacts
	// WHERE main_user = "mainUser value from request body"
	// AND contact_user = "contactUser value from request body"
	db.Table("contacts").
		Select("main_user, contact_user, is_blocked").
		Where("main_user = ? AND contact_user = ?", mainUser, contactUser).
		Scan(&result)
	if db.Error != nil {
		err := errors.New("Bad parsing")
		loger.Log.Errorf("Error has occurred: ", err)
		return true, err
	}
	switch result.IsBlocked {
	case 0:
		return false, nil
	case 1:
		return
	default:
		return
	}
	return
}
