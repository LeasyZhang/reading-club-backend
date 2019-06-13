package user

import (
	"fmt"
	"net/http"
	"reading-club-backend/database/entity"

	"github.com/jinzhu/gorm"

	"reading-club-backend/constant"
	"reading-club-backend/database"
	"reading-club-backend/dto"
)

var db *gorm.DB
var err error

// SaveOrUpdate save or update user entity
func SaveOrUpdate(user *entity.User) (tuser entity.User, userError *dto.UserErrorResponse) {

	db, err = database.GetDBConnection()
	if err != nil {
		return tuser, &dto.UserErrorResponse{ErrorCode: http.StatusInternalServerError, Error: err.Error()}
	}
	db.Save(tuser)

	return tuser, nil
}

// GetUserByName get user by user name
func GetUserByName(userName string) (entity.User, *dto.UserErrorResponse) {
	var userRsp entity.User
	var errorRsp dto.UserErrorResponse

	db, err = database.GetDBConnection()

	errors := db.Where("user_name = ?", userName).First(&userRsp).GetErrors()

	for _, err := range errors {
		if gorm.IsRecordNotFoundError(err) {
			errorRsp.Error = constant.UserNotFound
			errorRsp.ErrorCode = constant.UserNotFoundCode
			return userRsp, &errorRsp
		}
		errorRsp.Error = err.Error()
		errorRsp.ErrorCode = constant.InternalServerErrorCode
		return userRsp, &errorRsp
	}

	return userRsp, nil
}

// GetUserList : get user list
func GetUserList() (userList []entity.User, userError *dto.UserErrorResponse) {

	db, err := database.GetDBConnection()

	if err != nil {
		return userList, &dto.UserErrorResponse{ErrorCode: constant.CanNotConnectDatabaseCode, Error: err.Error()}
	}

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

	db, err := database.GetDBConnection()

	if err != nil {
		fmt.Println(err)
	}

	var user entity.User
	errors := db.Where("user_name = ?", username).First(&user).GetErrors()

	for _, err := range errors {
		fmt.Println(err)
		return
	}

	db.Delete(&user)
}
