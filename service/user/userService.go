package user

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"reading-club-backend/database"
)

var db *gorm.DB
var err error

// User Model
type User struct {
	ID           int    `gorm:"AUTO_INCREMENT;primary_key"`
	Username     string `uri:"name" binding:"required"`
	Email        string `uri:"email" binding:"required"`
	GroupName    string
	DonateStatus int
	DonateNumber int
	CreatedTime  time.Time
	UpdatedTime  time.Time
}

// TableName : table name for struct User
func (User) TableName() string {
	return "rc_user"
}

// AllUsers : List All Users
func AllUsers(c *gin.Context) {

	db, err = gorm.Open(database.DBEngine, database.DBName)
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

	db, err = gorm.Open(database.DBEngine, database.DBName)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	var newUser User
	if err := c.ShouldBindUri(&newUser); err != nil {
		c.JSON(400, gin.H{"msg": "parameter missing"})
		return
	}
	name := newUser.Username
	email := newUser.Email
	groupName := "DefaultGroup"
	donateStatus := 0
	donateNumber := 0
	createdTime := time.Now().UTC()
	updatedTime := time.Now().UTC()

	db.Create(&User{Username: name, Email: email, GroupName: groupName, DonateStatus: donateStatus, DonateNumber: donateNumber, CreatedTime: createdTime, UpdatedTime: updatedTime})

	c.JSON(200, gin.H{
		"message": "New User Created!",
	})
}

//DeleteUser : delete user of the given name
func DeleteUser(c *gin.Context) {
	db, err = gorm.Open(database.DBEngine, database.DBName)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	var username string
	username = c.Param("name")

	var user User
	db.Where("username = ?", username).Find(&user)

	if user.ID > 0 {
		db.Delete(&user)
	}

	c.JSON(200, gin.H{
		"message": "Delete User Successfully",
	})
}

//UpdateUser : update user by name
func UpdateUser(c *gin.Context) {
	db, err = gorm.Open(database.DBEngine, database.DBName)
	if err != nil {
		fmt.Println(err.Error())
	}
	
	defer db.Close()

	var updatedUser User
	if err := c.ShouldBindUri(&updatedUser); err != nil {
		c.JSON(400, gin.H{"msg": "parameter missing"})
		return
	}

	name := updatedUser.Username

	var user User
	db.Where("username = ?", name).Find(&user)

	if user.ID > 0 {
		user.Email = updatedUser.Email
		user.UpdatedTime = time.Now().UTC()
		db.Save(&user)
	}

	c.JSON(200, gin.H{
		"message": "Update User Successfully",
	})
}
