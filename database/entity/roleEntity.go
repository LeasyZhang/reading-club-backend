package entity

import (
	"time"
)

// Role Model
type Role struct {
	ID          int       `gorm:"AUTO_INCREMENT;primary_key"`
	UserName    string    `json:"username"`
	UserRole    string    `json:"userrole"`
	CreatedTime time.Time `json:"createdTime"`
	UpdatedTime time.Time `json:"updatedTime"`
}

// TableName role table name
func (Role) TableName() string {
	return "rc_user_role"
}
