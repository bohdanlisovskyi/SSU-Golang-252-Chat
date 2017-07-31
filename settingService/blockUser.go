package settingService

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"encoding/json"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"errors"
)

type BlockUserRequest struct {
	Header BlockUserRequestHeader `json:"header"`
	Body   BlockUserRequestBody `json:"body"`
}

type BlockUserRequestHeader struct {
	Type_    string `json:"type"`
	Command  string `json:"command"`
	UserName string `json:"userName"`
	Token    string `json:"token"`
}

type BlockUserRequestBody struct {
	IsBlocked   int `json:"is_blocked"`
	ContactName string `json:"name"`
}

func UnmarshalBlockUserRequest(byteRequest [] byte) (int, string, error) {
	var BlockUserRequest *BlockUserRequest
	err := json.Unmarshal(byteRequest, &BlockUserRequest)
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return -1, "", err
	}
	IsBlocked := BlockUserRequest.Body.IsBlocked
	ContactName := BlockUserRequest.Body.ContactName // check with contacts table
	if IsBlocked != 0 && IsBlocked != 1 {
		err := errors.New("IsBlocked value not valid")
		loger.Log.Errorf("IsBlocked value is not 1 or 0 :", err)
		return -1, "", err
	}
	return IsBlocked, ContactName, err
}

func BlockUnblockUser(byteRequest [] byte) (bool, error) {
	IsBlocked, ContactName, err := UnmarshalBlockUserRequest(byteRequest)
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	db, err := gorm.Open("sqlite3", "users.db/contacts")
	defer db.Close()
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	db.Update("IsBlocked", IsBlocked).Where("ContactName = ?", ContactName)
	if db.Error != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	return true, nil
}

// TODO func IsBlocked - checking if user is blocked
