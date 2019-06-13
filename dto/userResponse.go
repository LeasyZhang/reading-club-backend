package dto

import "time"

// UserResponse response of user api
type UserResponse struct {
	ID              int       `json:"id"`
	UserName        string    `json:"name"`
	Email          string    `json:"email"`
	GroupName      string    `json:"group"`
	DonateStatus   string    `json:"-"`
	DonateNumber   int    `json:"-"`
	CreatedTime    time.Time `json:"createdTime"`
	UpdatedTime    time.Time `json:"updatedTime"`
}

// UserErrorResponse error response of user api
type UserErrorResponse struct {
	ErrorCode int    `json:"errorCode"`
	Error     string `json:"error"`
}