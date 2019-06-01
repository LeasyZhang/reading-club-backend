package book

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	//use postgres database
	dbConn "reading-club-backend/database"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

// Book book entity
type Book struct {
	ID          int    `gorm:"AUTO_INCREMENT;primary_key"`
	BookName    string `json:"book_name"`
	Author      string `json:"author"`
	LeftAmount  int
	ISBN        string  `json:"isbn"`
	DoubanURL   string  `json:"douban_url"`
	ImageURL    string  `json:"image_url"`
	Price       float32 `json:"price"`
	Press       string  `json:"press"`
	CreatedTime time.Time
	UpdatedTime time.Time
}

// TableName related table name is rc_book
func (Book) TableName() string {
	return "rc_book"
}

// ViewBookDetail get book message by book id
func ViewBookDetail(c *gin.Context) {
	db, err = gorm.Open(dbConn.DBEngine, dbConn.DBName)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, gin.H{"error": "Database connection error"})
		return
	}

	defer db.Close()

	bookID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"error": "invalid parameter " + c.Param("id")})
		return
	}

	if bookID <= 0 {
		c.JSON(404, gin.H{"error": "the book you are looking for does not exist"})
		return
	}

	var book Book

	errors := db.First(&book, bookID).GetErrors()

	for _, err := range errors {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(404, gin.H{"error": "the book you are looking for does not exist"})
			return
		} else {
			c.JSON(500, gin.H{"error": err})
			return
		}
	}

	if book.ID <= 0 {
		c.JSON(404, gin.H{"error": "the book you are looking for does not exist"})
		return
	}

	c.JSON(200, gin.H{"message": book})
}

// FindBookByName get book message by book name
func FindBookByName(c *gin.Context) {
	db, err = gorm.Open(dbConn.DBEngine, dbConn.DBName)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, gin.H{"error": "Database connection error"})
		return
	}

	defer db.Close()

	bookName := c.Param("name")

	var book Book

	errors := db.Where("book_name = ?", bookName).First(&book).GetErrors()

	for _, err := range errors {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(404, gin.H{"error": "the book you are looking for does not exist"})
			return
		} else {
			c.JSON(500, gin.H{"error": err})
			return
		}
	}

	if book.ID <= 0 {
		c.JSON(404, gin.H{"error": "the book you are looking for does not exist"})
		return
	}

	c.JSON(200, gin.H{"message": book})
}
