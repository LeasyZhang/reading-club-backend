package book

import (
	"time"

	"github.com/jinzhu/gorm"

	"reading-club-backend/constant"
	bookDao "reading-club-backend/database/dao/book"
	historyDao "reading-club-backend/database/dao/history"
	userDao "reading-club-backend/database/dao/user"
	"reading-club-backend/database/entity"
	"reading-club-backend/dto"
)

var db *gorm.DB
var err error

// ViewBookDetail get book message by book id
func ViewBookDetail(bookID int) (bookRsp entity.Book, bookError *dto.BookErrorResponse) {

	bookRsp, bookError = bookDao.GetBookByID(bookID)

	if bookError != nil {
		return bookRsp, bookError
	}

	return bookRsp, nil
}

// FindBookByName get book message by book name
func FindBookByName(bookName string) (bookRsp entity.Book, bookError *dto.BookErrorResponse) {

	bookRsp, bookError = bookDao.GetBookByName(bookName)

	if bookError != nil {
		return bookRsp, bookError
	}

	return bookRsp, nil
}

// GetAllBooks Get All Books
func GetAllBooks() (bookList []entity.Book, err *dto.BookErrorResponse) {

	bookList, err = bookDao.GetAllBooks()

	if err != nil {
		return bookList, err
	}

	return bookList, nil
}

// BorrowBook borrow a book
func BorrowBook(userName string, bookID int) (bookRsp entity.Book, errorRsp *dto.BookErrorResponse) {

	//check book status(bookId)
	//update book left status(bookId)
	//add a history record(bookId, userId)
	//return the book user borrowed

	bookRsp, errorRsp = bookDao.GetBookByID(bookID)

	if errorRsp != nil {
		return bookRsp, errorRsp
	}

	currentUser, userError := userDao.GetUserByName(userName)

	if userError != nil {
		return bookRsp, &dto.BookErrorResponse{ErrorCode: userError.ErrorCode, Error: userError.Error}
	}

	if bookRsp.BookStatus == constant.BookBusy {
		return bookRsp, &dto.BookErrorResponse{ErrorCode: constant.BookAlreadyBorrowedCode, Error: constant.BookAlreadyBorrowed}
	}

	bookRsp.BookStatus = constant.BookBusy
	bookRsp.LeftAmount--
	bookRsp.UpdatedTime = time.Now()
	bookDao.SaveOrUpdate(&bookRsp)

	var bookHistory entity.BorrowHistory
	bookHistory.UserID = currentUser.ID
	bookHistory.UserName = userName
	bookHistory.BookID = bookID
	bookHistory.BookName = bookRsp.BookName
	bookHistory.BorrowDate = time.Now()
	bookHistory.HistoryStatus = constant.BookFree
	bookHistory.CreatedTime = time.Now()
	bookHistory.UpdatedTime = time.Now()
	historyDao.SaveOrUpdate(&bookHistory)

	return bookRsp, nil
}

// ReturnBook return a book
func ReturnBook(userName string, bookID int) (bookRsp entity.Book, errorRsp *dto.BookErrorResponse) {

	bookRsp, errorRsp = bookDao.GetBookByID(bookID)

	if errorRsp != nil {
		return bookRsp, errorRsp
	}

	_, userError := userDao.GetUserByName(userName)
	if userError != nil {
		return bookRsp, &dto.BookErrorResponse{ErrorCode: userError.ErrorCode, Error: userError.Error}
	}

	bookHistory, historyError := historyDao.GetUserBorrowedHistory(userName, bookID)
	if historyError != nil {
		return bookRsp, &dto.BookErrorResponse{ErrorCode: historyError.ErrorCode, Error: historyError.Error}
	}

	if bookHistory.HistoryStatus == constant.BookBusy {
		return bookRsp, &dto.BookErrorResponse{ErrorCode: constant.BookAlreadyBorrowedCode, Error: constant.BookAlreadyBorrowed}
	}

	bookRsp.BookStatus = constant.BookFree
	bookRsp.UpdatedTime = time.Now()
	bookRsp.LeftAmount++
	bookDao.SaveOrUpdate(&bookRsp)

	bookHistory.ReturnDate = time.Now()
	bookHistory.UpdatedTime = time.Now()
	bookHistory.HistoryStatus = constant.BookBusy
	historyDao.SaveOrUpdate(&bookHistory)

	return bookRsp, nil
}

// AddBook : add a new book
func AddBook(book entity.Book) entity.Book {

	bookDao.SaveOrUpdate(&book)

	return book
}

// UpdateBookInfo : update book information
func UpdateBookInfo(bookID int, description string) (bookRsp entity.Book, errorRsp *dto.BookErrorResponse) {

	bookRsp, errorRsp = bookDao.GetBookByID(bookID)

	if errorRsp != nil {
		return bookRsp, errorRsp
	}

	bookRsp.BookDescription = description
	bookRsp.UpdatedTime = time.Now()

	bookDao.SaveOrUpdate(&bookRsp)

	return bookRsp, nil
}

// DeleteBook : logic delete, just make it invisible
func DeleteBook(bookID int) {
	bookDao.DeleteBookByID(bookID)
}
