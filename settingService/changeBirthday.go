package settingService

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"errors"
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

	db, err := gorm.Open("sqlite3", "users.db/users") // change to common connection
	defer db.Close()// change to common connection
	if err != nil {// change to common connection
		loger.Log.Errorf("Error has occurred: ", err)// change to common connection
		return false, err// change to common connection
	}// change to common connection

	db.Update("birthday", birthday).Where("user_name = ?", userName)
	if db.Error != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	return true, nil
}
