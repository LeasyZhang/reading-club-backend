package user

import (
	"fmt"
	"reading-club-backend/database/entity"

	"github.com/jinzhu/gorm"

	"reading-club-backend/constant"
	"reading-club-backend/database"
	"reading-club-backend/dto"
)

var db *gorm.DB
var err error

// SaveOrUpdate save or update user entity
func SaveOrUpdate(user *entity.User) {

	db, err = database.GetDBConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Save(user)
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
