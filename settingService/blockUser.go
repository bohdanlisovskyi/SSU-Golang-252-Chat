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
		return
	}
	contactUser = body.ContactUser
	if contactUser == "" {
		err := errors.New("ContactUser value is empty")
		loger.Log.Error(err)
		return "", -1, err
	}
	isBlocked = body.IsBlocked // IsBlocked value could be int 0 or 1
	if isBlocked != 0 && isBlocked != 1 {
		err := errors.New("IsBlocked value is not valid. IsBlocked = " + string(isBlocked))
		loger.Log.Error(err)
		return "", -1, err
	}
	return contactUser, isBlocked, nil
}

// BlockUnblockUser perform blocking and unblocking of users by clicking the button on UI.
// Button works like a trigger: first push - block user, second push - unblock user.
func BlockUnblockUser(request *messageService.Message) (ok bool, err error) {
	contactUser, isBlocked, err := UnmarshalBlockUserRequestBody(request)
	if err != nil {
		loger.Log.Error(err)
		return
	}
	if request.Header.UserName == "" {
		err := errors.New("MainUser value is empty")
		loger.Log.Error(err)
		return false, err
	}
	db, err := database.GetStorage()
	if err != nil {
		return
	}
	contact := messageService.Contact{MainUser: request.Header.UserName, ContactUser: contactUser, IsBlocked: isBlocked}
	db.Model(&contact).Update("is_blocked", isBlocked)
	if db.Error != nil {
		loger.Log.Error(err)
		return
	}
	return true, nil
}

// IsUserBlocked checks if user is blocked in contacts table
func IsUserBlocked(request *messageService.Message) (isBlocked bool, err error) {
	isBlocked = true
	var body *messageService.MessageBody
	err = json.Unmarshal(request.Body, &body)
	if err != nil {
		loger.Log.Error(err)
		return
	}
	db, err := database.GetStorage()
	if err != nil {
		return
	}
	var result []messageService.Contact //variable for storing result of querying by contacts table
	var row messageService.Contact
	rows, err := db.Table("Contacts").
		Select("main_user, contact_user, is_blocked").
		Where("main_user = ? AND contact_user = ?", body.ReceiverName, request.Header.UserName).
		Rows()
	if err != nil {
		loger.Log.Errorf("Querying error:", err)
		return
	}
	for rows.Next() {
		if err := rows.Scan(&row.MainUser, &row.ContactUser, &row.IsBlocked); err != nil {
			loger.Log.Errorf("scan error:", err)
			return isBlocked, err
		}
		result = append(result, row)
	}
	switch row.IsBlocked {
	case 0:
		return false, nil
	case 1:
		return
	default:
		return
	}
	return
}
