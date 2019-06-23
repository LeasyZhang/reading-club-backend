package main

import (
	"github.com/gin-gonic/gin"

	"reading-club-backend/api"
	"reading-club-backend/database"
	"reading-club-backend/middleware"
)

func handleRequests() {
	router := gin.Default()

	router.Use(middleware.AllowCORS())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello~~"})
	})
	//user api
	router.GET("/user/list", api.GetUserList)
	router.POST("/user/add", api.AddUser)
	router.POST("/user/delete/:name", api.DeleteUser)
	router.POST("/user/update", api.UpdateUser)

	//book api
	router.GET("/book/detail/:ID", api.ViewBookDetail)
	router.GET("/book/search", api.SearchBook)
	router.GET("/book/list", api.GetAllBooks)
	router.POST("/book/add", api.AddBook)
	router.POST("/book/update", api.UpdateBook)
	router.POST("/book/delete/:bookID", api.DeleteBook)
	//borrow return api
	router.POST("/book/borrow", api.BorrowBook)
	router.POST("/book/return", api.ReturnBook)

	//history api
	router.GET("/history/user/:username", api.GetUserHistory)
	router.GET("/history/book/:bookID", api.GetBookHistory)

	//auth api
	loginHandler, _ := middleware.AuthMiddleWare()
	router.POST("/login", loginHandler.LoginHandler)
	router.POST("/logout", api.Logout)

	router.Run()
}

func main() {
	database.InitialDatabase()

	handleRequests()
}
