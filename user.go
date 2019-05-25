package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error
var dbName = "reading-club.db"

// User Model
type User struct {
	gorm.Model
	Name  string `uri:"name" binding:"required"`
	Email string `uri:"email" binding:"required"`
}

// InitialMigration : Database Init Operation, create database reading-club.db
func InitialMigration() {
	db, err = gorm.Open("sqlite3", dbName)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}

	defer db.Close()

	db.AutoMigrate(&User{})
}

// AllUsers : List All Users
func AllUsers(c *gin.Context) {

	db, err = gorm.Open("sqlite3", dbName)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	var users []User
	db.Find(&users)

	c.JSON(200, gin.H{
		"message": users,
	})
}

// NewUser : Add User
func NewUser(c *gin.Context) {

	db, err := gorm.Open("sqlite3", dbName)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	var newUser User
	if err := c.ShouldBindUri(&newUser); err != nil {
		c.JSON(400, gin.H{"msg": "parameter missing"})
		return
	}
	name := newUser.Name
	email := newUser.Email

	db.Create(&User{Name: name, Email: email})

	c.JSON(200, gin.H{
		"message": "New User Created!",
	})
}

//DeleteUser : delete user of the given name
func DeleteUser(c *gin.Context) {
	db, err := gorm.Open("sqlite3", dbName)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	var username string
	if err := c.ShouldBindUri(&username); err != nil {
		c.JSON(400, gin.H{"msg": "parameter missing"})
		return
	}
	name := username

	var user User
	db.Where("name = ?", name).Find(&user)

	if user.ID > 0 {
		db.Delete(&user)
	}
	
	c.JSON(200, gin.H{
		"message": "Delete User Successfully",
	})
}

//UpdateUser : update user by name
func UpdateUser(c *gin.Context) {
	db, err := gorm.Open("sqlite3", dbName)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	var updatedUser User
	if err := c.ShouldBindUri(&updatedUser); err != nil {
		c.JSON(400, gin.H{"msg": "parameter missing"})
		return
	}

	name := updatedUser.Name

	var user User
	db.Where("name = ?", name).Find(&user)
	
	if user.ID > 0 {
		user.Email = updatedUser.Email
		db.Save(&user)
	}
	
	c.JSON(200, gin.H{
		"message": "Update User Successfully",
	})
}
