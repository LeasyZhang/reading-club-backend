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
	//user api
	router.GET("/users", user.AllUsers)
	router.POST("/user/:name/:email", user.NewUser)
	router.DELETE("/user/:name", user.DeleteUser)
	router.PUT("/user/:name/:email", user.UpdateUser)

	//book api
	router.GET("/book/detail/:id", book.ViewBookDetail)
	router.GET("/search/book/:name", book.FindBookByName)
	router.GET("/books", book.GetAllBooks)
	router.POST("/book/add", book.AddBook)
	router.PUT("/book/update", book.UpdateBookInfo)
	router.DELETE("/book/delete/:bookId", book.DeleteBook)

	//history api
	router.GET("/history/user/:userId", history.GetUserBorrowHistory)
	router.GET("/history/book/:bookId", history.GetBookBorrowHistory)

	//borrow return api
	router.POST("/borrow", book.BorrowBook)
	router.POST("/return", book.ReturnBook)

	router.Run()
}

func main() {
	database.InitialDatabase()

	handleRequests()
}
