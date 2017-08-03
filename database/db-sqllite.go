package database

import (
	"sync"

	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/8tomat8/SSU-Golang-252-Chat/server/config"
	"github.com/jinzhu/gorm"
)

var (
	once sync.Once
	db   *gorm.DB
)

func GetStorage() (*gorm.DB, error) {
	var err error

	once.Do(func() {
		settings := config.GetConfig()
		db, err = gorm.Open(settings.Storage.Driver, settings.Storage.Name)
		if err != nil {
			loger.Log.Errorf("ERROR connection to SqlLite3 %s", err.Error())
		}
	})
	return db, err
}
