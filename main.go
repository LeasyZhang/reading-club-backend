package main

import (
	"github.com/gin-gonic/gin"

	"reading-club-backend/service/book"
	"reading-club-backend/service/user"
	"reading-club-backend/database"
	"reading-club-backend/middleware"
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

	router.Run()
}

func main() {
	database.InitialDatabase()

	handleRequests()
}
