package settingService

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"encoding/json"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"errors"
)

// BlockUserRequest is a structure for request to block user
type BlockUserRequest struct {
	Header RequestHeader `json:"header"`
	Body   json.RawMessage `json:"body"`
}

// UnmarshalBlockUserRequest function unmarshal request for blocking of users into map[string]interface{}
func UnmarshalBlockUserRequest(byteRequest [] byte) (map[string]interface{}, error) {
	var BlockUserRequest map[string]interface{}
	err := json.Unmarshal(byteRequest, BlockUserRequest)
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return nil, err
	}
	IsBlocked := BlockUserRequest["is_blocked"]
	if IsBlocked != 0 && IsBlocked != 1 {
		err := errors.New("IsBlocked value not valid")
		loger.Log.Errorf("IsBlocked value is not 1 or 0 :", err)
		return nil, err
	}
	return BlockUserRequest, nil
}

// BlockUnblockUser perform blocking and unblocking of users
func BlockUnblockUser(byteRequest [] byte) (bool, error) {
	BlockUserRequest, err := UnmarshalBlockUserRequest(byteRequest)
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	IsBlocked := BlockUserRequest["is_blocked"]
	ContactName := BlockUserRequest["user_name"]
	if IsBlocked == nil || ContactName == nil {
		loger.Log.Errorf("Some field or fields are empty")
		return false, err
	}

	db, err := gorm.Open("sqlite3", "users.db/contacts") // change to common connection
	defer db.Close()                                     // change to common connection
	if err != nil { // change to common connection
		loger.Log.Errorf("Error has occurred: ", err) // change to common connection
		return false, err                             // change to common connection
	} // change to common connection

	db.Update("is_blocked", IsBlocked).Where("user_name = ?", ContactName) // change according to DB data
	if db.Error != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	return true, nil
}

// TODO func IsBlocked - checking if user is blocked
