package settingService

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

// ChangePassRequest is a structure for request to change password
type ChangePassRequest struct {
	Header RequestHeader `json:"header"`
	Body   json.RawMessage `json:"body"`
}

// UnmarshalChangePassRequest is a function for unmarshaling request (from [] byte JSON into map structure)
func UnmarshalChangePassRequest(byteRequest [] byte) (map[string]interface{}, error) {
	var unmarshaledRequest map[string]interface{}
	err := json.Unmarshal(byteRequest, &unmarshaledRequest)
	if err != nil {
		loger.Log.Errorf("Error has occured: ", err)
		return nil, err
	}
	return unmarshaledRequest, err
}

func ChangePass(changePassRequest [] byte) (bool, error) {
	unmarshaledRequest, err := UnmarshalChangePassRequest(changePassRequest)
	if err != nil {
		loger.Log.Errorf("Error has occured: ", err)
		return nil, err
	}
	userName := unmarshaledRequest["user_name"]
	OldPass := unmarshaledRequest["old_pass"]
	NewPass := unmarshaledRequest["new_pass"]
	if OldPass == NewPass {
		loger.Log.Warn("Old and new passwords are the same ")
	}
	db, err := gorm.Open("sqlite3", "users.db/users") // change to common connection
	defer db.Close()
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	db.Update("password", NewPass).Where("user_name = ?", userName)
	if db.Error != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	return true, nil
}
