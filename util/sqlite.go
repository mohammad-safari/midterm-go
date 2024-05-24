package util

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToSQLite() (*gorm.DB, error) {
	var db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
