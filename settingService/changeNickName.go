package settingService

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

// ChangeNickNameRequest is a structure for request to change nick-name
type ChangeNickNameRequest struct {
	Header RequestHeader `json:"header"`
	Body   json.RawMessage `json:"body"`
}

// UnmarshalChangeNickNameRequest is a function for unmarshaling request (from [] byte JSON to map)
func UnmarshalChangeNickNameRequest(changeNickNameRequest [] byte) (map[string]interface{}, error) {
	var unmarshaledRequest map[string]interface{}
	err := json.Unmarshal(changeNickNameRequest, unmarshaledRequest)
	if err != nil {
		loger.Log.Errorf("Error has occured: ", err)
		return nil, err
	}
	return unmarshaledRequest, nil
}

func ChangeNickName(changeNickNameRequest [] byte) (bool, error) {
	unmarshaledRequest, err := UnmarshalChangeNickNameRequest(changeNickNameRequest)
	if err != nil {
		loger.Log.Errorf("Error has occured: ", err)
		return nil, err
	}
	nickName := unmarshaledRequest["nick_name"]
	userName := unmarshaledRequest["user_name"]
	db, err := gorm.Open("sqlite3", "users.db/users") // change to common connection
	defer db.Close()
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	db.Update("nick_name", nickName).Where("user_name = ?", userName)
	if db.Error != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	return true, nil
}
