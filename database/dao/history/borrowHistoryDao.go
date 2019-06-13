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
	var errorRsp dto.HistoryErrorResponse

	db, err = database.GetDBConnection()

	if err != nil {
		errorRsp.Error = constant.CanNotConnectDatabase
		errorRsp.ErrorCode = constant.CanNotConnectDatabaseCode
		return bookHistory, &errorRsp
	}

	bookErros := db.Where("user_name = ? and book_id = ? and history_status = ?", userName, bookID, constant.BookFree).First(&bookHistory).GetErrors()

	for _, err := range bookErros {
		if gorm.IsRecordNotFoundError(err) {
			errorRsp.Error = constant.HistoryNotFound
			errorRsp.ErrorCode = constant.HistoryNotFoundCode
			return bookHistory, &errorRsp
		}

		errorRsp.Error = constant.InternalServerError
		errorRsp.ErrorCode = constant.InternalServerErrorCode
		return bookHistory, &errorRsp
	}

	return bookHistory, nil
}
