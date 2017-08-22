package settingService

import (
	"encoding/json"
	"errors"

	"github.com/8tomat8/SSU-Golang-252-Chat/database"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
)

// BlockUserRequestBody is a custom body for BlockUserRequest
type BlockUserRequestBody struct {
	ContactUser string `json:"contact_user"`
	IsBlocked   int `json:"is_blocked"`
}

// UnmarshalBlockUserRequestBody function unmarshals request for blocking of users and retrieves values of
// contactUser and isBlocked variables.
func UnmarshalBlockUserRequestBody(request *messageService.Message) (contactUser string, isBlocked int, err error) {
	var body *BlockUserRequestBody
	err = json.Unmarshal(request.Body, &body)
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return "", -1, err
	}
	contactUser = body.ContactUser
	isBlocked = body.IsBlocked // IsBlocked value could be int 0 or 1
	if isBlocked != 0 && isBlocked != 1 {
		err := errors.New("IsBlocked value is not valid. IsBlocked = " + string(isBlocked))
		loger.Log.Errorf("IsBlocked value is not 1 or 0:", err)
		return "", -1, err
	}
	if contactUser == "" {
		err := errors.New("ContactUser value is empty")
		loger.Log.Errorf("Error has occurred: ", err)
		return "", -1, err
	}
	return contactUser, isBlocked, nil
}

// BlockUnblockUser perform blocking and unblocking of users by clicking the button on UI.
// Button works like a trigger: first push - block user, second push - unblock user.
func BlockUnblockUser(request *messageService.Message) (bool, error) {
	contactUser, isBlocked, err := UnmarshalBlockUserRequestBody(request)
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	mainUser := request.Header.UserName
	if mainUser == "" {
		err := errors.New("MainUser value is empty")
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	db, err := database.GetStorage() // common gorm-connection from database package
	if err != nil {
		loger.Log.Errorf("DB error has occurred: ", err)
		return false, err
	}
	contact := messageService.Contact{MainUser: mainUser, ContactUser: contactUser, IsBlocked: isBlocked}
	db.Model(&contact).Update("is_blocked", isBlocked)
	if db.Error != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	return true, nil
}

// IsUserBlocked checks if user is blocked in contacts table
func IsUserBlocked(request *messageService.Message) (isBlocked bool, err error) {
	isBlocked = true
	var body *messageService.MessageBody
	err = json.Unmarshal(request.Body, &body)
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return true, err
	}
	mainUser := body.ReceiverName
	contactUser := request.Header.UserName
	db, err := database.GetStorage()
	if err != nil {
		loger.Log.Errorf("DB error has occurred: ", err)
		return true, err
	}
	var result *messageService.Contact //variable for storing result of querying by contacts table
	// SELECT main_user, contact_user, is_blocked FROM contacts
	// WHERE main_user = "mainUser value from request body"
	// AND contact_user = "contactUser value from request header"
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
