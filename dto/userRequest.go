package dto

//UserRequest request body of user api
type UserRequest struct {
	UserName	string	`json:"userName"`
	Email	string	`json:"email"`
}