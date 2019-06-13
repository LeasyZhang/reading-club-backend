package entity

import "time"

//BorrowHistory borrow list
type BorrowHistory struct {
	ID            int `gorm:"AUTO_INCREMENT;primary_key"`
	UserID        int
	BookID        int
	UserName      string
	BookName      string
	BorrowDate    time.Time
	ReturnDate    time.Time
	DueDate       time.Time
	HistoryStatus int
	CreatedTime   time.Time
	UpdatedTime   time.Time
}

// TableName history table name
func (BorrowHistory) TableName() string {
	return "rc_book_history"
}