package dto

import "time"

//LoginResponse response of login request
type LoginResponse struct {
	AccessToken	string	`json:"accessToken"`
	ExpiredTime	time.Time	`json:"-"`
}