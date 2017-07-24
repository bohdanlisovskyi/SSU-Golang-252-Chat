package auth

import (
	"database/sql"
	"fmt"
_   "github.com/mattn/go-sqlite3"
	"github.com/Greckas/SSU-Golang-252-Chat/messageService"
	"log"
	"github.com/Greckas/SSU-Golang-252-Chat/loger"
)

var (
db, _ = sql.Open("go-sqlite3", "messager.db")
)
//request map from server



func parse_request(body_data [] byte) (Login, Password, Name string, err error)  {
	request_data, err := messageService.UnmarshalRequest(body_data)
	if err != nil {
		log.Println("invalid recived data", err)
		return nil, nil, nil, err

	}

	Login := request_data["login"]
	Password := request_data["password"]
	Name := request_data["name"]
	return Login, Password, Name, nil
}

func retrive_db_data(body_data [] byte) (bool, error){
	login,_,_, err := parse_request(body_data)
	if err != nil {
		loger.Log.Errorf("Parsing from server have failed", err)
	}
	rows, err := db.Query("SELECT login FROM users WHERE login ="+ login)
	if rows != nil{
		loger.Log.Errorf("User with this login already exists", err)
		return false, err
	}
	return true, nil
}

type User struct {
Name      string
Login     string
Password  string
}


func (u *User) InsertIntoUsers(body_data[] byte) (bool, error) {
	ok, _ := retrive_db_data(body_data)
	tx, _ := db.Begin()
	defer db.Close()
	if ok {

		stmp, _ := tx.Prepare("insert into users(name, login, password) values (?, ?, ?)")

		_, err := stmp.Exec(u.Name, u.Login, u.Password)
		// перепитати як мають записуватись дані в базу корректно

		if err != nil {
			loger.Log.Errorf("query didn`t execute", err)
			return false, err
		}

		tx.Commit()
		return true, err
	} else {
		tx.Rollback()

	}
	return true, nil
}




//func main() {
//	InsertIntoUsers(body_data[] byte)
//
//}