package database

import (
	"sync"
	"database/sql"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/server/config"
)

var SqlLite *sql.DB

func init() {

	once.Do(GetSqlLiteStorage)
}
var once sync.Once

func GetSqlLiteStorage() {

	settings := config.GetConfig()

	db, err := sql.Open(settings.Storage[0].Driver, settings.Storage[0].Name)

	if err == nil {

		loger.Log.Errorf("Error connection to SqlLite3 %s", err.Error())

		return
	}

	SqlLite = db
}