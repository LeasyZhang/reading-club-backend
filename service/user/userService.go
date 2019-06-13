package user

import (
	"time"

	userDao "reading-club-backend/database/dao/user"
	"reading-club-backend/database/entity"
	"reading-club-backend/dto"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// AllUsers : List All Users
func AllUsers() (userList []entity.User, userError *dto.UserErrorResponse) {

	userList, userError = userDao.GetUserList()

	return userList, userError
}

// NewUser : Add User
func NewUser(username string, email string) (user entity.User, userError *dto.UserErrorResponse) {

	var newUser entity.User

	newUser.Username = username
	newUser.Email = email
	newUser.GroupName = "Integration"
	newUser.DonateStatus = 0
	newUser.DonateNumber = 0
	newUser.CreatedTime = time.Now().UTC()
	newUser.UpdatedTime = time.Now().UTC()

	user, userError = userDao.SaveOrUpdate(&newUser)

	return newUser, userError
}

//DeleteUser : delete user of the given name
func DeleteUser(username string) {
	userDao.DeleteUser(username)
}

//UpdateUser : update user by name
func UpdateUser(email string, username string) (user entity.User, userError *dto.UserErrorResponse) {

	user, userError = userDao.GetUserByName(username)
	if userError != nil {
		return user, userError
	}

	user.Email = email
	user, userError = userDao.SaveOrUpdate(&user)
	return user, userError
}
