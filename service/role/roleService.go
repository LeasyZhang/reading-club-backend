package user

import (
	roleDao "reading-club-backend/database/dao/role"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// GetUserRole : find role by a given name
func GetUserRole(username string) string {
	return roleDao.GetRoleByName(username)
}
