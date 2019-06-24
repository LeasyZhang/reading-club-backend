package database

import (
	"os"

	"github.com/jinzhu/gorm"
	//postgres database
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Conn database connection
var Conn *gorm.DB

//TestDBName test database connection(local database)
var testDBName = "host=localhost port=5432 user=joe.zhang dbname=mydb password=19950209 sslmode=disable"
var prodDBName = os.Getenv("DATABASE_URL")

//DBName production : os.Getenv("DATABASE_URL") local: TestDBName
var DBName = prodDBName

//DBEngine engine name for database connection
var DBEngine = "postgres"

//GetDBConnection open database connection
func GetDBConnection() (*gorm.DB, error) {
	db, err := gorm.Open(DBEngine, DBName)
	if err != nil {
		panic("Failed to connect to database " + err.Error())
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(20)

	Conn = db
	return db, err
}
