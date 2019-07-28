package role

import (
	"reading-club-backend/database/entity"

	"github.com/jinzhu/gorm"

	db "reading-club-backend/database"
	"reading-club-backend/dto"
)

var err error

// SaveOrUpdate save or update user entity
func SaveOrUpdate(role *entity.Role) (entity.Role, *dto.UserErrorResponse) {

	db.Conn.Save(role)

	return *role, nil
}

// GetRoleByName get user by user name
func GetRoleByName(userName string) string {
	var role entity.Role
	errors := db.Conn.Where("user_name = ?", userName).First(&role).GetErrors()

	for _, err := range errors {
		if gorm.IsRecordNotFoundError(err) {
			return "illegal"
		}
		return "illegal"
	}

	return role.UserRole
}
