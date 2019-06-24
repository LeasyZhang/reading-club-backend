package database

import (
	"os"

	"github.com/jinzhu/gorm"
	//postgres database
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Conn database connection
var Conn *gorm.DB
var err error

//TestDBName test database connection(local database)
var testDBName = "host=localhost port=5432 user=joe.zhang dbname=mydb password=19950209 sslmode=disable"
var prodDBName = os.Getenv("DATABASE_URL")

//DBName production : os.Getenv("DATABASE_URL") local: TestDBName
var DBName = prodDBName

//DBEngine engine name for database connection
var DBEngine = "postgres"

//GetDBConnection open database connection
func GetDBConnection() (*gorm.DB, error) {
	Conn, err = gorm.Open(DBEngine, DBName)
	if err != nil {
		panic("Failed to connect to database " + err.Error())
	}
	Conn.DB().SetMaxIdleConns(20)
	Conn.DB().SetMaxOpenConns(100)
	Conn.DB().Ping()
	return Conn, err
}
