package history

import (
	"fmt"
	"net/http"
	"time"

	"reading-club-backend/database"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	//use postgres database

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

//BorrowHistory borrow list
type BorrowHistory struct {
	ID            int `gorm:"AUTO_INCREMENT;primary_key"`
	UserID        int
	BookID        int
	UserName      string
	BookName      string
	BorrowDate    time.Time
	ReturnDate    time.Time
	DueDate       time.Time
	HistoryStatus int
	CreatedTime   time.Time
	UpdatedTime   time.Time
}

// TableName history table name
func (BorrowHistory) TableName() string {
	return "rc_book_history"
}

// GetUserBorrowHistory get current user's borrow list
func GetUserBorrowHistory(c *gin.Context) {

	db, err = gorm.Open(database.DBEngine, database.DBName)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, gin.H{"error": "Database connection error"})
		return
	}

	defer db.Close()

	userID := c.Param("userId")

	var historyList []BorrowHistory

	errors := db.Where("user_id = ?", userID).Find(&historyList).GetErrors()

	for _, err := range errors {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": historyList})
}

// GetBookBorrowHistory get current book's borrow history
func GetBookBorrowHistory(c *gin.Context) {

	db, err = gorm.Open(database.DBEngine, database.DBName)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, gin.H{"error": "Database connection error"})
		return
	}

	defer db.Close()

	bookID := c.Param("bookId")

	var historyList []BorrowHistory

	errors := db.Where("book_id = ?", bookID).Find(&historyList).GetErrors()

	for _, err := range errors {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": historyList})
}
