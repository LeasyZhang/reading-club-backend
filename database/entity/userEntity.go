package entity

import (
	"time"
)

// User Model
type User struct {
	ID           int `gorm:"AUTO_INCREMENT;primary_key"`
	Username     string
	Password     string `json:"-"`
	Email        string
	GroupName    string
	DonateStatus int
	DonateNumber int
	CreatedTime  time.Time
	UpdatedTime  time.Time
}

// TableName history table name
func (User) TableName() string {
	return "rc_user"
}
