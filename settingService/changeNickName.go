package settingService

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"errors"
)

// ChangeNickNameRequest is a structure for request to change nick-name of users
type ChangeNickNameRequest struct {
	Header RequestHeader `json:"header"`
	Body   json.RawMessage `json:"body"`
}

// UnmarshalChangeNickNameRequest is a function for unmarshaling request for changing nick-name from []byte JSON
// to map[string]interface{}
func UnmarshalChangeNickNameRequest(changeNickNameRequest [] byte) (map[string]interface{}, error) {
	var unmarshaledRequest map[string]interface{}
	err := json.Unmarshal(changeNickNameRequest, unmarshaledRequest)
	if err != nil {
		loger.Log.Errorf("Error has occured: ", err)
		return nil, err
	}
	return unmarshaledRequest, nil
}

// ChangeNickName perform changing nick-name of users
func ChangeNickName(changeNickNameRequest [] byte) (bool, error) {
	unmarshaledRequest, err := UnmarshalChangeNickNameRequest(changeNickNameRequest)
	if err != nil {
		loger.Log.Errorf("Error has occured: ", err)
		return false, err
	}
	nickName := unmarshaledRequest["nick_name"]
	userName := unmarshaledRequest["user_name"]
	if nickName == nil || userName == nil{
		err := errors.New("Empty field or fields")
		loger.Log.Errorf("Some field or fields are empty: ")
		return false, err
	}

	db, err := gorm.Open("sqlite3", "users.db/users") // change to common connection
	defer db.Close()// change to common connection
	if err != nil {// change to common connection
		loger.Log.Errorf("Error has occurred: ", err)// change to common connection
		return false, err// change to common connection
	}// change to common connection

	db.Update("nick_name", nickName).Where("user_name = ?", userName)
	if db.Error != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	return true, nil
}
