package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func handleRequests() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{"msg": "hello, user function is now ready"})
	})
	router.GET("/users", AllUsers)
	router.POST("/user/:name/:email", NewUser)
	router.DELETE("/user/:name", DeleteUser)
	router.PUT("/user/:name/:email", UpdateUser)
	
	router.Run(":8080")
}

func main() {
	fmt.Println("Go Orm Turorial")

	InitialMigration()

	handleRequests()
}
