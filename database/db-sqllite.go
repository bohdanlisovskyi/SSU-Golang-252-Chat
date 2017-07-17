package database

import (
	"database/sql"
	"github.com/8tomat8/SSU-Golang-252-Chat/config"
	"github.com/8tomat8/SSU-Golang-252-Chat/chat-log"
)

type StorageInterface struct {
	SqlLite map[bool]*sql.DB
}

var storage StorageInterface

func init() {

	storage = StorageInterface{}

	storage.SqlLite = make(map[bool]*sql.DB)
}

func GetSqlLiteStorage() *sql.DB {

	if store, exists := storage.SqlLite[true]; exists {
		return store
	}

	settings := config.GetConfig()

	db, err := sql.Open(settings.Storage.Driver, settings.Storage.Name)

	chatlog.HandleError(err, "Connect Error: ")

	storage.SqlLite[true] = db

	return storage.SqlLite[true]
}