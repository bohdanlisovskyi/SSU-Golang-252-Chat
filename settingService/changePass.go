package settingService

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

// ChangePassRequest is a structure for request to change password
type ChangePassRequest struct {
	Header RequestHeader `json:"header"`
	Body   ChangePassRequestBody `json:"body"`
}

// ChangePassRequestBody is a structure for passing new and old passwords
type ChangePassRequestBody struct {
	OldPass string
	NewPass string
}

// UnmarshalChangePassRequest is a function for unmarshaling request (from [] byte JSON to ChangePassRequest structure)
func UnmarshalChangePassRequest(changePassRequest [] byte) (*ChangePassRequest, error) {
	var RequestStructure *ChangePassRequest
	err := json.Unmarshal(changePassRequest, &RequestStructure)
	if err != nil {
		loger.Log.Errorf("Error has occured: ", err)
		return RequestStructure, err
	}
	return RequestStructure, nil
}

func ChangePass(changePassRequest [] byte) (bool, error) {
	RequestStructure, err := UnmarshalChangePassRequest(changePassRequest)
	if err != nil {
		loger.Log.Errorf("Error has occured: ", err)
		return nil, err
	}
	userName := RequestStructure.Header.UserName
	OldPass := RequestStructure.Body.OldPass
	NewPass := RequestStructure.Body.NewPass
	if OldPass == NewPass {
		loger.Log.Warn("Old and new passwords are the same ")
	}
	db, err := gorm.Open("sqlite3", "users.db")
	defer db.Close()
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	db.Update("NewPass", NewPass).Where("user_name = ?", userName)
	if db.Error != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	return true, nil
}
