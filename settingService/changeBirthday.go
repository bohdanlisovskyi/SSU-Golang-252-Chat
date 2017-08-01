package settingService

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

// ChangeBirthdayRequest is a structure for request to change nick-name
type ChangeBirthdayRequest struct {
	Header RequestHeader `json:"header"`
	Body   json.RawMessage `json:"body"`
}

// UnmarshalBirthdayRequest is a function for unmarshaling request (from [] byte JSON to map)
func UnmarshalChangeBirthdayRequest(changeBirthdayRequest [] byte) (map[string]interface{}, error) {
	var unmarshaledRequest map[string]interface{}
	err := json.Unmarshal(changeBirthdayRequest, unmarshaledRequest)
	if err != nil {
		loger.Log.Errorf("Error has occured: ", err)
		return nil, err
	}
	return unmarshaledRequest, nil
}

func ChangeBirthday(changeBirthdayRequest [] byte) (bool, error) {
	unmarshaledRequest, err := UnmarshalChangeBirthdayRequest(changeBirthdayRequest)
	if err != nil {
		loger.Log.Errorf("Error has occured: ", err)
		return nil, err
	}
	birthday := unmarshaledRequest["birthday"]
	userName := unmarshaledRequest["user_name"]
	db, err := gorm.Open("sqlite3", "users.db/users") // change to common connection
	defer db.Close()
	if err != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	db.Update("birthday", birthday).Where("user_name = ?", userName)
	if db.Error != nil {
		loger.Log.Errorf("Error has occurred: ", err)
		return false, err
	}
	return true, nil
}
