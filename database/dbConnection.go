package database

import (
	"os"

	"github.com/jinzhu/gorm"
	//postgres database
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"reading-club-backend/config"
)

//Conn database connection
var Conn *gorm.DB

//TestDBName test database connection(local database)
var prodDBName = os.Getenv("DATABASE_URL")
var dbConnectionURL = config.Configuration.DB.URL
var maxIdleConns = config.Configuration.DB.MaxIdleConnections
var maxOpenConns = config.Configuration.DB.MaxOpenConnections

//DBEngine engine name for database connection
var dBEngine = config.Configuration.DB.DbEngine

//GetDBConnection open database connection
func GetDBConnection() (*gorm.DB, error) {
	db, err := gorm.Open(dBEngine, dbConnectionURL)
	if err != nil {
		panic("Failed to connect to database " + err.Error())
	}
	db.DB().SetMaxIdleConns(maxIdleConns)
	db.DB().SetMaxOpenConns(maxOpenConns)

	Conn = db
	return db, err
}
