package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"

	"github.com/8tomat8/SSU-Golang-252-Chat/database"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/messageService"
	"github.com/8tomat8/SSU-Golang-252-Chat/server/customers"
	"golang.org/x/crypto/bcrypt"
)

type clientResult struct {
	main_user    string
	contact_user string `gorm:"contact_user"`
	is_blocked   int
}

type nickName struct {
	nick_name string `gorm:"contact_user"`
}

func getUserByName(UserName string) (*messageService.Authentification, error) {

	db, err := database.GetStorage()
	if err != nil {
		loger.Log.Errorf("Failed to open db", err)
		return nil, err
	}
	db_search := db.Where("user_name=?", UserName)
	ret := &messageService.Authentification{}
	err = db_search.First(ret).Error
	if err != nil {
		loger.Log.Errorf("Failed to find user in DB")
		return nil, err
	}
	return ret, err
}

func Login(username, password string) (*messageService.Authentification, string, error) {
	foundUser, err := getUserByName(username)
	if err != nil {
		loger.Log.Errorf("No user with that Username")
		return nil, "", err
	}

	// Generate "hash" to check from user password
	hash, err := bcrypt.GenerateFromPassword([]byte(foundUser.Password), bcrypt.DefaultCost)
	if err != nil {
		loger.Log.Errorf("hash generation failed!", err)
		return nil, "", err
	}
	// Comparing the password with the hash
	if err := bcrypt.CompareHashAndPassword(hash, []byte(password)); err != nil {
		loger.Log.Errorf("Password check failed!", err)
		return nil, "", err
	}

	token := randToken()
	return foundUser, token, nil
}

func VerifyToken(message *messageService.Message, client customers.Client) (bool, error) {
	clientToken := message.Header.Token
	serverToken := client.Token
	if clientToken == "" {
		err := errors.New("Empty token, pls relogin!")
		return false, err
	}
	if serverToken == "" {
		err := errors.New("No token in connection!")
		return false, err
	}
	if clientToken == serverToken {
		return true, nil
	}
	err := errors.New("token verification failed!")
	return false, err

}

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func SendNickName(message *messageService.Authentification) (string, error) {
	db, err := database.GetStorage()
	if err != nil {
		loger.Log.Errorf("Failed to open db andsend contacts", err)
		return "", err
	}
	authentification := &messageService.Authentification{}
	db.Where("user_name = ?", message.UserName).First(&authentification)
	return authentification.NickName, nil
}
func SendContacts(cont *messageService.Message) (*messageService.ClientContacts, error) {
	db, err := database.GetStorage()
	if err != nil {
		loger.Log.Errorf("Failed to open db andsend contacts", err)
		return nil, err
	}
	var result []clientResult

	rows, err := db.Table("Contacts").Select("main_user, contact_user, is_blocked").
		Where("main_user = ? and is_blocked == 0", cont.Header.UserName).Rows()
	if err != nil {
		loger.Log.Errorf("query from contacts error:", err)
		return nil, err
	}
	var row clientResult
	for rows.Next() {
		if err := rows.Scan(&row.main_user, &row.contact_user, &row.is_blocked); err != nil {
			loger.Log.Errorf("scan error:", err)
			return nil, err
		}
		result = append(result, row)
	}
	var nick []string

	for _, item := range result {
		var nickRow nickName
		rows, err := db.Table("Authentifications").Select("nick_name").
			Where("user_name = ?", item.contact_user).Rows()
		if err != nil {
			loger.Log.Errorf("query from auth error:", err)
			return nil, err
		}
		for rows.Next() {
			if err := rows.Scan(&nickRow.nick_name); err != nil {
				loger.Log.Errorf("scan error:", err)
				return nil, err
			}
			nick = append(nick, nickRow.nick_name)
		}
	}
	var contact messageService.ClientContact
	contact = messageService.ClientContact{}
	var contacts *messageService.ClientContacts
	contacts = &messageService.ClientContacts{}
	for i, j := range result {
		contact.NickName = nick[i]
		contact.UserName = j.contact_user
		contact.IsBlocked = j.is_blocked
		contacts.ContactsList = append(contacts.ContactsList, contact)
	}
	return contacts, err
}
