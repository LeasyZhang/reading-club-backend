package entity

import (
	"time"
)

// User Model
type User struct {
	ID           int       `gorm:"AUTO_INCREMENT;primary_key"`
	UserName     string    `json:"username"`
	Password     string    `json:"-"`
	Email        string    `json:"email"`
	GroupName    string    `json:"group"`
	DonateStatus int       `json:"-"`
	DonateNumber int       `json:"-"`
	CreatedTime  time.Time `json:"createdTime"`
	UpdatedTime  time.Time `json:"updatedTime"`
}

// TableName user table name
func (User) TableName() string {
	return "rc_user"
}
