package settingService

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"encoding/json"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"errors"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
)

// BlockUserRequest is a structure for request to block user
type BlockUserRequest struct {
	Header messageService.MessageHeader `json:"header"`
	Body   json.RawMessage `json:"body"`
}

// BlockUserRequestBody is a custom body for BlockUserRequest
type BlockUserRequestBody struct {
	MainUser string `json:"main_user"`
	ContactUser string `json:"contact_user"`
	IsBlocked int `json:"is_blocked"`
}

// UnmarshalBlockUserRequestBody function unmarshal request for blocking of users into BlockUserRequestBody struct,
// check value of is_blocked variable, second return parameter - IsBlocked value
func UnmarshalBlockUserRequestBody(byteRequest messageService.Message) (*BlockUserRequestBody, int, error) {
	var body *BlockUserRequestBody
	err := json.Unmarshal(byteRequest.Body, &body)
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return nil, nil, err
	}
	IsBlocked := body.IsBlocked // IsBlocked value could be int 0 or 1
	if IsBlocked != 0 && IsBlocked != 1 {
		err := errors.New("IsBlocked value not valid")
		loger.Log.Errorf("IsBlocked value is not 1 or 0 :", err)
		return nil, nil, err
	}
	return body, IsBlocked, nil
}

// BlockUnblockUser perform blocking and unblocking of users by clicking the button on UI.
// Button works like a trigger: first push - user is blocked, second push - user is unblocked
func BlockUnblockUser(request BlockUserRequest) (bool, error) {
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
	db, err := GetStorage() // common gorm-connection from database package
	defer db.Close()
	if err != nil {
		loger.Log.Errorf("DB error has occurred: ", err)
		return false, err
	}
	// UPDATE contacts SET IsBlocked= "IsBlocked value" WHERE main_user = "mainUser value" AND  contact_user = "contactUser value"
	db.Model(&Contact).Where("main_user = ? AND contact_user = ?", mainUser, contactUser).Update("is_blocked", IsBlocked)
	if db.Error != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	return true, nil
}

// IsUserBlocked checks if user is blocked for chatting in contacts table
func IsUserBlocked(byteRequest [] byte) (isBlocked int, err error) {
	request, err := messageService.UnmarshalMessage(byteRequest)
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return nil, err
	}
	contactUser := request.Body.ReceiverName
	mainUser := request.Header.UserName
	db, err := GetStorage() // common gorm-connection from database package
	defer db.Close()
	if err != nil {
		loger.Log.Errorf("DB error has occurred: ", err)
		return nil, err
	}

	values := db.Where("main_user = ? AND contact_user = ?", mainUser, contactUser).Find(&contacts).values
	//// SELECT * FROM contacts WHERE main_user = 'mainUser' AND contact_user = 'contactUser'

	isBlocked = values["is_blocked"]
	if isBlocked == 0 {
		return 0, nil
	}
	if isBlocked == 1 {
		return 1, nil
	}
	parsingErr := errors.New("Bad parsing")
	return nil, parsingErr
}
