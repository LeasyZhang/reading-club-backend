package database

import (
	"os"

	"github.com/jinzhu/gorm"
	//postgres database
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var database *gorm.DB
var err error

//TestDBName test database connection(local database)
var testDBName = "host=localhost port=5432 user=joe.zhang dbname=mydb password=19950209 sslmode=disable"
var prodDBName = os.Getenv("DATABASE_URL")

//DBName production : os.Getenv("DATABASE_URL") local: TestDBName
var DBName = prodDBName

//DBEngine engine name for database connection
var DBEngine = "postgres"

//InitialDatabase open database connection
func InitialDatabase() {
	database, err = gorm.Open(DBEngine, DBName)
	if err != nil {
		panic("Failed to connect to database " + err.Error())
	}

	defer database.Close()

	database.DB().Ping()
}