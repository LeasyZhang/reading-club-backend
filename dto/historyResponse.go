package dto

import "time"

// HistoryResponse response of history api
type HistoryResponse struct {
	ID              int       `json:"id"`
	BookName        string    `json:"name"`
	Author          string    `json:"author"`
	ISBN            string    `json:"ISBN"`
	DoubanURL       string    `json:"doubanUrl"`
	ImageURL        string    `json:"imageUrl"`
	Price           float32   `json:"price"`
	Press           string    `json:"press"`
	BookDescription string    `json:"description"`
	CreatedTime     time.Time `json:"createdTime"`
	UpdatedTime     time.Time `json:"updatedTime"`
}

// HistoryErrorResponse error response of history api
type HistoryErrorResponse struct {
	ErrorCode int    `json:"errorCode"`
	Error     string `json:"error"`
}