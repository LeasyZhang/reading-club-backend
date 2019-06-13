package api

import (
	"net/http"
	historyService "reading-club-backend/service/history"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetUserHistory borrow history of user
func GetUserHistory(c *gin.Context) {

	username := c.Param("username")
	historyList, err := historyService.GetUserBorrowHistory(username)

	if err != nil {
		c.JSON(err.ErrorCode, gin.H{"error": err.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": historyList})
	return
}

//GetBookHistory borrow history of a book
func GetBookHistory(c *gin.Context) {

	bookID, err := strconv.Atoi(c.Param("bookID"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	historyList, historyErr := historyService.GetBookBorrowHistory(bookID)

	if historyErr != nil {
		c.JSON(historyErr.ErrorCode, gin.H{"error": historyErr.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": historyList})
	return
}
