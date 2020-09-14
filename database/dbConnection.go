package database

import (
	"github.com/jinzhu/gorm"
	"os"
	//postgres database
	"reading-club-backend/config"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"reading-club-backend/database/entity"
)

//Conn database connection
var Conn *gorm.DB

//GetDBConnection open database connection
func GetDBConnection() (*gorm.DB, error) {
	var dbConnectionURL string
	dbConnectionURL, exists := os.LookupEnv("DATABASE_URL")
	if !exists {
		dbConnectionURL = config.Configuration.DB.URL
	}

	maxIdleConns := config.Configuration.DB.MaxIdleConnections
	maxOpenConns := config.Configuration.DB.MaxOpenConnections

	//DBEngine engine name for database connection
	dBEngine := config.Configuration.DB.DbEngine

	db, err := gorm.Open(dBEngine, dbConnectionURL)
	if err != nil {
		panic("Failed to connect to database " + err.Error())
	}
	db.DB().SetMaxIdleConns(maxIdleConns)
	db.DB().SetMaxOpenConns(maxOpenConns)

	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Book{})
	db.AutoMigrate(&entity.Role{})
	db.AutoMigrate(&entity.Feature{})
	db.AutoMigrate(&entity.BorrowHistory{})

	Conn = db

	return db, err
}
