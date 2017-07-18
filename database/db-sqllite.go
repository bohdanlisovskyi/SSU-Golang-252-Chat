package database

import (
	"sync"
	"database/sql"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/server/config"
)

var (
	once sync.Once
	sqlLite *sql.DB
)

func GetStorage() *sql.DB {

	once.Do(func() {

		settings := config.GetConfig()

		db, err := sql.Open(settings.Storage.Driver, settings.Storage.Name)

		if err == nil {

			loger.Log.Errorf("Error connection to SqlLite3 %s", err.Error())

			return
		}

		sqlLite = db
	})

	return sqlLite
}