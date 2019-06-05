package main

import (
	"github.com/gin-gonic/gin"

	"reading-club-backend/database"
	"reading-club-backend/middleware"
	"reading-club-backend/service/book"
	"reading-club-backend/service/history"
	"reading-club-backend/service/user"
)

func handleRequests() {
	router := gin.Default()

	router.Use(middleware.AllowCORS())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "hello, user function is now ready"})
	})
	router.GET("/users", user.AllUsers)
	router.POST("/user/:name/:email", user.NewUser)
	router.DELETE("/user/:name", user.DeleteUser)
	router.PUT("/user/:name/:email", user.UpdateUser)
	router.GET("/book/:id", book.ViewBookDetail)
	router.GET("/search/book/:name", book.FindBookByName)
	router.GET("/borrow/book/:bookId", book.BorrowBook)
	router.GET("/return/book/:bookId", book.ReturnBook)
	router.GET("/history/user/:userId", history.GetUserBorrowHistory)
	router.GET("/books", book.BookList)
	router.GET("/history/book/:bookId", history.GetBookBorrowHistory)

	router.Run()
}

func main() {
	database.InitialDatabase()

	handleRequests()
}
