package history

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	//use postgres database

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

//BorrowHistory borrow list
type BorrowHistory struct {
	ID          int `gorm:"AUTO_INCREMENT;primary_key"`
	UserID      int
	BookID      int
	CreatedTime time.Time
	UpdatedTime time.Time
}

// GetUserBorrowHistory get current user's borrow list
func GetUserBorrowHistory(c *gin.Context) {

	userID := c.Param("userId")

	var historyList BorrowHistory

	fmt.Println(userID)

	c.JSON(http.StatusAccepted, gin.H{"message": historyList})
}
