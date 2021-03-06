package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"reading-club-backend/api"
	"reading-club-backend/config"
	"reading-club-backend/database"
	"reading-club-backend/middleware"
)

func handleRequests() {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	router.Use(middleware.AllowCORS())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello~~"})
	})

	router.GET("/user/list", api.GetUserList)

	router.OPTIONS("/user/list", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})
	router.OPTIONS("/user/add", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})
	router.OPTIONS("/user/delete/:name", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})
	router.OPTIONS("/user/add/update", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	//book api
	router.GET("/book/detail/:ID", api.ViewBookDetail)
	router.GET("/book/search", api.SearchBook)
	router.GET("/book/list", api.GetAllBooks)

	router.OPTIONS("/book/detail/:ID", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})
	router.OPTIONS("/book/search", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})
	router.OPTIONS("/book/list", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})
	router.OPTIONS("/book/add", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})
	router.OPTIONS("/book/update", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})
	router.OPTIONS("/book/delete/:bookID", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})
	router.OPTIONS("/book/borrow", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})
	router.OPTIONS("/book/return", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})
	router.OPTIONS("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})
	router.OPTIONS("/logout", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	//history api
	router.GET("/history/user/:username", api.GetUserHistory)
	router.GET("/history/book/:bookID", api.GetBookHistory)

	//auth api
	loginHandler, _ := middleware.AuthMiddleWare()
	router.POST("/login", loginHandler.LoginHandler)
	router.POST("/logout", api.Logout)
	router.POST("/heroku-deployed", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "successfully deployed"})
	})
	router.Use(loginHandler.MiddlewareFunc())
	{
		router.POST("/feature/add/:name", api.AddFeature)
		router.POST("/book/add", api.AddBook)
		router.POST("/book/update", api.UpdateBook)
		router.POST("/book/delete/:bookID", api.DeleteBook)
		router.POST("/feature/enable/:name", api.EnableFeature)
		router.POST("/feature/disable/:name", api.DisableFeature)

		//borrow return api
		router.POST("/book/borrow", api.BorrowBook)
		router.POST("/book/return", api.ReturnBook)
		router.POST("/user/add", api.AddUser)
		router.POST("/user/delete/:name", api.DeleteUser)
		router.POST("/user/update", api.UpdateUser)
		router.GET("/refresh_token", loginHandler.RefreshHandler)
	}

	router.Run()
}

func main() {
	config.InitConfiguration()
	db, err := database.GetDBConnection()

	if err != nil {
		panic("fatal error, failed to connect to database")
	}
	defer db.Close()

	handleRequests()
}
