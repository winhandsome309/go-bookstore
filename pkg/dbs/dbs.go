package dbs

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// const (
// 	DSN = "host=localhost user=postgres password=30092002 dbname=bookstore port=5432 sslmode=disable"
// )

var (
	db *gorm.DB
)

func Connect(databaseURI string) {
	d, err := gorm.Open(postgres.Open(databaseURI), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
