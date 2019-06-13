package main

import (
	"github.com/gin-gonic/gin"

	"reading-club-backend/api"
	"reading-club-backend/database"
	"reading-club-backend/middleware"
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
	router.GET("/user/list", user.AllUsers)
	router.POST("/user/add", user.NewUser)
	router.DELETE("/user/delete/:name", user.DeleteUser)
	router.POST("/user/update", user.UpdateUser)

	//book api
	router.GET("/book/detail/:ID", api.ViewBookDetail)
	router.GET("/book/search", api.SearchBook)
	router.GET("/book/list", api.GetAllBooks)
	router.POST("/book/add", api.AddBook)
	router.POST("/book/update", api.UpdateBook)
	router.DELETE("/book/delete/:bookID", api.DeleteBook)
	//borrow return api
	router.POST("/book/borrow", api.BorrowBook)
	router.POST("/book/return", api.ReturnBook)

	//history api
	router.GET("/history/user/:userName", history.GetUserBorrowHistory)
	router.GET("/history/book/:bookID", history.GetBookBorrowHistory)

	//auth api
	router.POST("/login", api.Login)
	router.POST("/logout", api.Logout)

	router.Run()
}

func main() {
	database.InitialDatabase()

	handleRequests()
}
