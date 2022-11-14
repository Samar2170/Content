package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {

	var err error
	db, err = gorm.Open(postgres.Open(DBURI), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Index{}, &IndexData{})
	db.AutoMigrate(&ForexPair{}, &ForexData{})
}
