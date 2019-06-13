package dto

// HistoryRequest borrow/return book request body
type HistoryRequest struct {

	UserName	string	`json:"username"`
	BookID	int	`json:"bookID"`
}