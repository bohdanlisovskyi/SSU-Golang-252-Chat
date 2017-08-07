package database

import (
	"sync"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/server/config"
)

var (
	once sync.Once
	sqlLite *gorm.DB
	err error
)

func GetStorage() (*gorm.DB, error) {

	once.Do(func() {

		settings := config.GetConfig()

		db, err := gorm.Open(settings.Storage.Driver, settings.Storage.Name)

		if err == nil {

			loger.Log.Errorf("Error connection to SqlLite3 %s", err.Error())

			return
		}

		sqlLite = db
	})

	return sqlLite, err
}