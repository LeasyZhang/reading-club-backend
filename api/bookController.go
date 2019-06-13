package api

import (
	"net/http"
	"reading-club-backend/constant"
	"reading-club-backend/database/entity"
	"reading-club-backend/dto"
	bookService "reading-club-backend/service/book"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// ViewBookDetail get book detail by unique id
func ViewBookDetail(c *gin.Context) {

	bookID, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, bookErr := bookService.ViewBookDetail(bookID)
	if bookErr != nil {
		c.JSON(bookErr.ErrorCode, gin.H{"error": bookErr.Error})
		return
	}

	c.JSON(http.StatusOK, book)
	return
}

// SearchBook search book by name
func SearchBook(c *gin.Context) {

	bookName := c.Query("name")

	book, err := bookService.FindBookByName(bookName)

	if err != nil {
		c.JSON(err.ErrorCode, gin.H{"error": err.Error})
		return
	}

	c.JSON(http.StatusOK, book)
	return
}

// GetAllBooks get all visible books
func GetAllBooks(c *gin.Context) {

	bookList, err := bookService.GetAllBooks()

	if err != nil {
		c.JSON(err.ErrorCode, gin.H{"error": err.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": bookList})
	return
}

// AddBook : add a book
func AddBook(c *gin.Context) {

	var bookRequest dto.BookRequest
	err := c.BindJSON(&bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var book entity.Book
	book.BookName = bookRequest.Name
	book.Author = bookRequest.Author
	book.LeftAmount = 1
	book.BookStatus = constant.BookFree
	book.ISBN = bookRequest.ISBN
	book.DoubanURL = bookRequest.DoubanURL
	book.ImageURL = bookRequest.ImageURL
	book.Price = bookRequest.Price
	book.Press = bookRequest.Press
	book.BookDescription = bookRequest.Description
	book.Visibility = 1
	book.CreatedTime = time.Now()
	book.UpdatedTime = time.Now()

	book = bookService.AddBook(book)
	c.JSON(http.StatusOK, book)
	return
}

// UpdateBook : update book description
func UpdateBook(c *gin.Context) {
	var bookRequest dto.BookRequest

	err := c.BindJSON(&bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, bookErr := bookService.UpdateBookInfo(bookRequest.ID, bookRequest.Description)

	if bookErr != nil {
		c.JSON(bookErr.ErrorCode, gin.H{"error": bookErr.Error})
		return
	}

	c.JSON(http.StatusOK, book)
	return
}

// DeleteBook : delete a book by id
func DeleteBook(c *gin.Context) {

	bookID, err := strconv.Atoi(c.Param("bookID"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookService.DeleteBook(bookID)

	c.JSON(http.StatusOK, gin.H{"id": bookID})
	return
}

// BorrowBook : borrow a book
func BorrowBook(c *gin.Context) {

	var borrowRequest dto.HistoryRequest
	err := c.ShouldBindJSON(&borrowRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, bookErr := bookService.BorrowBook(borrowRequest.UserName, borrowRequest.BookID)

	if bookErr != nil {
		c.JSON(bookErr.ErrorCode, gin.H{"error": bookErr.Error})
		return
	}

	c.JSON(http.StatusOK, book)
	return
}

// ReturnBook : return a book
func ReturnBook(c *gin.Context) {

	var returnRequest dto.HistoryRequest
	err := c.ShouldBindJSON(&returnRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, bookErr := bookService.ReturnBook(returnRequest.UserName, returnRequest.BookID)

	if bookErr != nil {
		c.JSON(bookErr.ErrorCode, gin.H{"error": bookErr.Error})
		return
	}

	c.JSON(http.StatusOK, book)
	return
}
