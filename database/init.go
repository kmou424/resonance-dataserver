package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var Conn *gorm.DB

func InitSQLite() {
	var err error
	Conn, err = gorm.Open(sqlite.Open("file:ignore.db?cache=shared&mode=memory"), &gorm.Config{
		IgnoreRelationshipsWhenMigrating: true,
	})
	if err != nil {
		log.Fatal("failed to connect database", "error", err)
	}

	migrate()
	importData()
}
