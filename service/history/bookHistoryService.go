package history

import (
	historyDao "reading-club-backend/database/dao/history"
	"reading-club-backend/database/entity"
	"reading-club-backend/dto"
)

// GetUserBorrowHistory get current user's borrow list
func GetUserBorrowHistory(username string) (historyList []entity.BorrowHistory, historyErr *dto.HistoryErrorResponse) {

	historyList, err := historyDao.GetUserHistory(username)

	return historyList, err
}

// GetBookBorrowHistory get current book's borrow history
func GetBookBorrowHistory(bookID int) (historyList []entity.BorrowHistory, historyErr *dto.HistoryErrorResponse) {

	historyList, err := historyDao.GetBookHistory(bookID)

	return historyList, err
}
