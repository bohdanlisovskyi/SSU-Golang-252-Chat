package settingService

import (
	"encoding/json"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"errors"
	"github.com/8tomat8/SSU-Golang-252-Chat/database"

)

// ChangeBirthdayRequest is a structure for request to change user nick-name
type ChangeBirthdayRequest struct {
	Header RequestHeader `json:"header"`
	Body   json.RawMessage `json:"body"`
}

// UnmarshalBirthdayRequest is a function for unmarshaling request for changing nick-name from []byte JSON
// to map[string]interface{}
func UnmarshalChangeBirthdayRequest(changeBirthdayRequest [] byte) (map[string]interface{}, error) {
	var unmarshaledRequest map[string]interface{}
	err := json.Unmarshal(changeBirthdayRequest, unmarshaledRequest)
	if err != nil {
		loger.Log.Errorf("Error has occured: ", err)
		return nil, err
	}
	return unmarshaledRequest, nil
}

// ChangeBirthday perform changing birthday of user
func ChangeBirthday(changeBirthdayRequest [] byte) (bool, error) {
	unmarshaledRequest, err := UnmarshalChangeBirthdayRequest(changeBirthdayRequest)
	if err != nil {
		loger.Log.Errorf("Error has occured: ", err)
		return false, err
	}
	birthday := unmarshaledRequest["birthday"]
	userName := unmarshaledRequest["user_name"]
	if birthday == nil || userName == nil{
		err := errors.New("Empty field or fields")
		loger.Log.Errorf("Some field or fields are empty: ")
		return false, err
	}
	db, err := GetStorage() // common gorm-connection from database package
	defer db.Close()
	if err != nil {
		loger.Log.Errorf("DB error has occurred: ", err)
		return false, err
	}
	db.Update("birthday", birthday).Where("user_name = ?", userName)
	if db.Error != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	return true, nil
}
