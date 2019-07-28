package test

import (
	"os"
	"reading-club-backend/database"
	roleService "reading-club-backend/service/role"
	"testing"
)

func TestMain(m *testing.M) {
	// add a another comment
	db, err := database.GetDBConnection()

	if err != nil {
		panic("fatal error, failed to connect to database")
	}
	defer db.Close()

	os.Exit(m.Run())
}

func TestGetRoleByUserName(t *testing.T) {
	userName := "joe"
	roleName := roleService.GetUserRole(userName)

	if roleName != "user" {
		t.Error()
	}

	userName = "nobody"
	roleName = roleService.GetUserRole(userName)
	if roleName != "illegal" {
		t.Error()
	}
}
