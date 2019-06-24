package user

import (
	"fmt"
	"reading-club-backend/database/entity"

	"github.com/jinzhu/gorm"

	"reading-club-backend/constant"
	"reading-club-backend/database"
	"reading-club-backend/dto"
	"reading-club-backend/util"
)

var db = database.Conn
var err error

// SaveOrUpdate save or update user entity
func SaveOrUpdate(user *entity.User) (tuser entity.User, userError *dto.UserErrorResponse) {

	db.Save(tuser)

	return tuser, nil
}

// GetUserByName get user by user name
func GetUserByName(userName string) (entity.User, *dto.UserErrorResponse) {
	var userRsp entity.User
	errors := db.Where("user_name = ?", userName).First(&userRsp).GetErrors()

	for _, err := range errors {
		if gorm.IsRecordNotFoundError(err) {
			return userRsp, &dto.UserErrorResponse{Error: constant.UserNotFound, ErrorCode: constant.UserNotFoundCode}
		}
		return userRsp, &dto.UserErrorResponse{Error: err.Error(), ErrorCode: constant.InternalServerErrorCode}
	}

	return userRsp, nil
}

// GetUserList : get user list
func GetUserList() (userList []entity.User, userError *dto.UserErrorResponse) {

	errors := db.Find(&userList).GetErrors()

	for _, err := range errors {
		if gorm.IsRecordNotFoundError(err) {
			return userList, &dto.UserErrorResponse{ErrorCode: constant.UserNotFoundCode, Error: constant.UserNotFound}
		}
		return userList, &dto.UserErrorResponse{ErrorCode: constant.InternalServerErrorCode, Error: err.Error()}
	}

	return userList, nil
}

//DeleteUser : delete user information from database
func DeleteUser(username string) {

	var user entity.User
	errors := db.Where("user_name = ?", username).First(&user).GetErrors()

	for _, err := range errors {
		fmt.Println(err)
		return
	}

	db.Delete(&user)
}

// FindUserByNameAndPassword : find user by given name and password
func FindUserByNameAndPassword(username string, password string) bool {

	encoded := util.Encrypt(password)

	var user entity.User
	errors := db.Where("user_name ~* ? and password = ?", username, encoded).Find(&user).GetErrors()
	for _, err := range errors {
		fmt.Println(err)
		return false
	}

	return true
}
