package database

import (
	"sync"

	"github.com/Greckas/SSU-Golang-252-Chat/loger"
	"github.com/Greckas/SSU-Golang-252-Chat/server/config"
	"github.com/jinzhu/gorm"
)

var (
	once    sync.Once
	sqlLite *gorm.DB
)

func GetStorage() (*gorm.DB, error) {
	var err error

	once.Do(func() {
		settings := config.GetConfig()
		sqlLite, err = gorm.Open(settings.Storage.Driver, settings.Storage.Name)
		if err != nil {
			loger.Log.Errorf("ERROR connection to SqlLite3 %s", err.Error())
		}
	})
	return sqlLite, err
}
