package contactService

import (
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"encoding/json"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	dbGorm "github.com/8tomat8/SSU-Golang-252-Chat/database"
)

type Authentifications struct {
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Password string
}

type Users struct {
	UserName  string `json:"username"`
	Birthday  int    `json:"birthday"`
	AboutUser string `json:"aboutuser"`
	UserImage string `json:"userimage`
}

type Contacts struct {
	MainUser    string `json:"mainuser"`
	ContactUser string `json:"contactuser"`
	IsBlocked   int    `json:"isblocked`
}

type Search struct {
	SearchContact string `json:"searchvalue"`
}

func AddContact(contact Contacts) (bool, error){
	db, err := dbGorm.GetStorage()
	if err != nil{
		loger.Log.Errorf("AddContact db connection ERROR: ", err)
		return false, err
	}

	err = db.Create(&contact).Error
	if err != nil {
		loger.Log.Errorf("AddContact: Insert new record into contacts table ERROR: ", err)
		return false, err
	}
	return true, err
}

func AddNewUser(user Users) (bool, error){
	db, err := dbGorm.GetStorage()
	if err != nil{
		loger.Log.Errorf("AddNewUser db connection ERROR: ", err)
		return false, err
	}

	err = db.Create(&user).Error
	if err != nil {
		loger.Log.Errorf("AddNewUser: Insert new record into users table ERROR: ", err)
		return false, err
	}
	return true, err
}

func GetContacts(id string) ([]Users, error){
	users := []Users{}
	db, err := dbGorm.GetStorage()
	if err != nil{
		loger.Log.Errorf("GetContacts db connection ERROR: ", err)
		return users, err
	}

	err = db.Joins("inner join contacts on contacts.contact_user = users.user_name").Where("is_blocked == 0 and contacts.main_user = ?", id).Find(&users).Error
	if err != nil {
		loger.Log.Errorf("GetContacts func ERROR: ", err)
	}
	return users, err
}

func GetUserInfo(id string) (Users, error){
	var user Users
	db, err := dbGorm.GetStorage()
	if err != nil{
		loger.Log.Errorf("GetUserInfo db connection ERROR: ", err)
		return user, err
	}

	err = db.First(&user, "user_name = ?", id).Error
	if err != nil {
		loger.Log.Errorf("GetUserInfo func ERROR: ", err)
	}
	return user, err
}

func RemoveContacts(id string)(bool, error){
	db, err := dbGorm.GetStorage()
	if err != nil{
		loger.Log.Errorf("RemoveContacts db connection ERROR: ", err)
		return false, err
	}

	tr := db.Begin()
	err = tr.Where("contact_user = ?", id).Delete(Contacts{}).Error
	if err != nil {
		loger.Log.Errorf("Delete contacts from contacts table ERROR: ", err)
		tr.Rollback()
		return false, err
	}

	tr.Commit()
	return true, err
}

func RemoveUser(id string)(bool, error) {
	db, err := dbGorm.GetStorage()
	if err != nil{
		loger.Log.Errorf("RemoveUser db connection ERROR: ", err)
		return false, err
	}

	// begin a transaction
	tr := db.Begin()

	if _, err := RemoveContacts(id); err != nil {
		loger.Log.Errorf("Delete contacts ERROR: ", err)
		tr.Rollback()
		return false, err
	}

	err = tr.Where("main_user = ?", id).Delete(Contacts{}).Error
	if err != nil {
		loger.Log.Errorf("Delete record from contacts table ERROR: ", err)
		tr.Rollback()
		return false, err
	}

	err = tr.Where("user_name = ?", id).Delete(Users{}).Error
	if err != nil {
		loger.Log.Errorf("Delete record from users table ERROR: ", err)
		tr.Rollback()
		return false, err
	}

	err = tr.Where("user_name = ?", id).Delete(Authentifications{}).Error
	if err != nil {
		loger.Log.Errorf("Delete record from authentifications table ERROR: ", err)
		tr.Rollback()
		return false, err
	}
	tr.Commit()
	return true, err
}

func SearchContacts(search string)([]Users, error){
	users := []Users{}
	db, err := dbGorm.GetStorage()
	if err != nil{
		loger.Log.Errorf("SearchContact db connection ERROR: ", err)
		return users, err
	}

	err = db.Raw("SELECT u.* FROM authentifications as a inner join users as u on a.user_name = u.user_name WHERE a.user_name LIKE ? or a.nick_name LIKE ?",
		"%" + search + "%", "%" + search + "%").Scan(&users).Error
	if err != nil {
		loger.Log.Errorf("SearchContact func search data ERROR: ", err)
	}
	return users, err
}

func UnmarshalContacts(request messageService.Message) (*Contacts, error) {
	var body *Contacts
	err := json.Unmarshal(request.Body, &body)
	if err != nil {
		loger.Log.Errorf("UnmarshalContacts func Error has occurred: ", err)
		return body, err
	}
	return body, err
}

func UnmarshalUsers(request messageService.Message) (*Users, error) {
	var body *Users
	err := json.Unmarshal(request.Body, &body)
	if err != nil {
		loger.Log.Errorf("UnmarshalUsers func Error has occurred: ", err)
		return body, err
	}
	return body, err
}

func UnmarshalSearch(request messageService.Message) (*Search, error) {
	var body *Search
	err := json.Unmarshal(request.Body, &body)
	if err != nil {
		loger.Log.Errorf("UnmarshalSearch func Error has occurred: ", err)
		return body, err
	}
	return body, err
}