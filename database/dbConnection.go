package database

import (
	"os"

	"github.com/jinzhu/gorm"

	//postgres database
	"reading-club-backend/config"

	"reading-club-backend/database/entity"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Conn database connection
var Conn *gorm.DB

//GetDBConnection open database connection
func GetDBConnection() (*gorm.DB, error) {
	dbConnectionURL := os.Getenv("DATABASE_URL")
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
