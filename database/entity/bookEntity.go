package entity

import (
	"time"
)
// Book book entity
type Book struct {
	ID              int    `gorm:"AUTO_INCREMENT;primary_key" json:"id"`
	BookName        string `json:"name"`
	Author          string `json:"author"`
	LeftAmount      int		`json:"-"`
	BookStatus      int     `json:"status"`
	ISBN            string  `json:"ISBN"`
	DoubanURL       string  `json:"doubanUrl"`
	ImageURL        string  `json:"imageUrl"`
	Price           float32 `json:"price"`
	Press           string  `json:"press"`
	BookDescription string  `json:"description"`
	Visibility      int		`json:"-"`
	CreatedTime     time.Time	`json:"createdTime"`
	UpdatedTime     time.Time	`json:"updatedTime"`
}

// TableName table name of book entity
func (Book) TableName() string {
	return "rc_book"
}