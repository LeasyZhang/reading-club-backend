package history

import (
	"fmt"
	"reading-club-backend/database/entity"

	"github.com/jinzhu/gorm"

	"reading-club-backend/constant"
	"reading-club-backend/database"
	"reading-club-backend/dto"
)

var db *gorm.DB
var err error

// SaveOrUpdate save or update borrow history entity
func SaveOrUpdate(history *entity.BorrowHistory) {
	db, err = database.GetDBConnection()

	if err != nil {
		fmt.Println(err)
		return
	}

	db.Save(history)
}

// GetUserBorrowedHistory get user borrow history by username bookid
func GetUserBorrowedHistory(userName string, bookID int) (entity.BorrowHistory, *dto.HistoryErrorResponse) {

	var bookHistory entity.BorrowHistory

	db, err = database.GetDBConnection()

	if err != nil {
		return bookHistory, &dto.HistoryErrorResponse{ErrorCode: constant.CanNotConnectDatabaseCode, Error: constant.CanNotConnectDatabase}
	}

	bookErros := db.Where("user_name = ? and book_id = ? and history_status = ?", userName, bookID, constant.BookFree).First(&bookHistory).GetErrors()

	for _, err := range bookErros {
		if gorm.IsRecordNotFoundError(err) {
			return bookHistory, &dto.HistoryErrorResponse{ErrorCode: constant.HistoryNotFoundCode, Error: constant.HistoryNotFound}
		}

		return bookHistory, &dto.HistoryErrorResponse{ErrorCode: constant.InternalServerErrorCode, Error: constant.InternalServerError}
	}

	return bookHistory, nil
}

// GetUserHistory : get user borrow history by username
func GetUserHistory(username string) (historyList []entity.BorrowHistory, historyErr *dto.HistoryErrorResponse) {

	db, err = database.GetDBConnection()

	if err != nil {
		return historyList, &dto.HistoryErrorResponse{ErrorCode: constant.CanNotConnectDatabaseCode, Error: err.Error()}
	}

	errors := db.Where("user_name = ?", username).Find(&historyList).GetErrors()

	for _, err := range errors {
		if gorm.IsRecordNotFoundError(err) {
			return historyList, &dto.HistoryErrorResponse{ErrorCode: constant.HistoryNotFoundCode, Error: constant.HistoryNotFound}
		}
		return historyList, &dto.HistoryErrorResponse{ErrorCode: constant.InternalServerErrorCode, Error: constant.InternalServerError}
	}

	return historyList, nil
}

// GetBookHistory : get book borrow history by bookID
func GetBookHistory(bookID int) (historyList []entity.BorrowHistory, historyErr *dto.HistoryErrorResponse) {

	db, err = database.GetDBConnection()

	if err != nil {
		return historyList, &dto.HistoryErrorResponse{ErrorCode: constant.CanNotConnectDatabaseCode, Error: err.Error()}
	}

	errors := db.Where("book_id = ?", bookID).Find(&historyList).GetErrors()

	for _, err := range errors {
		if gorm.IsRecordNotFoundError(err) {
			return historyList, &dto.HistoryErrorResponse{ErrorCode: constant.HistoryNotFoundCode, Error: constant.HistoryNotFound}
		}
		return historyList, &dto.HistoryErrorResponse{ErrorCode: constant.InternalServerErrorCode, Error: constant.InternalServerError}
	}

	return historyList, nil
}
