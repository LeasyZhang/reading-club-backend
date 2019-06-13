package api

import (
	"net/http"
	"reading-club-backend/dto"
	userService "reading-club-backend/service/user"

	"github.com/gin-gonic/gin"
)

//Login : user login
func Login(c *gin.Context) {
	c.JSON(http.StatusOK, dto.LoginResponse{AccessToken: "ewiuyrofqigoqwiurhf137946q8fwghwioufhq"})
}

// Logout : user logout
func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Logout Success!"})
}

// AddUser : add a user
func AddUser(c *gin.Context) {

	var userRequest dto.UserRequest
	err := c.ShouldBindJSON(&userRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser, userError := userService.NewUser(userRequest.UserName, userRequest.Email)

	if userError != nil {
		c.JSON(userError.ErrorCode, gin.H{"error": userError.Error})
		return
	}

	c.JSON(http.StatusOK, newUser)
}

// UpdateUser : update user email
func UpdateUser(c *gin.Context) {

	var userRequest dto.UserRequest
	err := c.ShouldBindJSON(&userRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, userError := userService.UpdateUser(userRequest.Email, userRequest.UserName)

	if userError != nil {
		c.JSON(userError.ErrorCode, gin.H{"error": userError.Error})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// GetUserList : get user list
func GetUserList(c *gin.Context) {
	userList, userError := userService.AllUsers()
	if userError != nil {
		c.JSON(userError.ErrorCode, gin.H{"error": userError.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": userList})
}

// DeleteUser : delete user by given name
func DeleteUser(c *gin.Context) {
	userName := c.Param("username")

	userService.DeleteUser(userName)

	c.JSON(http.StatusOK, gin.H{"message": "Delete Success!"})
}
