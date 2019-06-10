package book

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	//use postgres database
	"reading-club-backend/constant"
	"reading-club-backend/database"
	dbConn "reading-club-backend/database"
	"reading-club-backend/service/history"
	"reading-club-backend/service/user"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

// Book book entity
type Book struct {
	ID              int    `gorm:"AUTO_INCREMENT;primary_key"`
	BookName        string `json:"book_name"`
	Author          string `json:"author"`
	LeftAmount      int
	BookStatus      int     `json:"book_status"`
	ISBN            string  `json:"isbn"`
	DoubanURL       string  `json:"douban_url"`
	ImageURL        string  `json:"image_url"`
	Price           float32 `json:"price"`
	Press           string  `json:"press"`
	BookDescription string  `json:"book_description"`
	Visibility      int
	CreatedTime     time.Time
	UpdatedTime     time.Time
}

// UserRequest request body
type UserRequest struct {
	UserID int `json:"userId"`
	BookID int `json:"bookId"`
}

// BookRequest book request body
type BookRequest struct {
	BookID      int     `json:"bookID"`
	BookName    string  `json:"bookName"`
	Author      string  `json:"author"`
	BookStatus  int     `json:"bookStatus"`
	ISBN        string  `json:"isbn"`
	DoubanURL   string  `json:"doubanUrl"`
	ImageURL    string  `json:"imageUrl"`
	Price       float32 `json:"price"`
	Press       string  `json:"press"`
	Description string  `json:"description"`
}

// TableName related table name is rc_book
func (Book) TableName() string {
	return "rc_book"
}

// ViewBookDetail get book message by book id
func ViewBookDetail(c *gin.Context) {
	db, err = gorm.Open(dbConn.DBEngine, dbConn.DBName)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, gin.H{"error": "Database connection error"})
		return
	}

	defer db.Close()

	bookID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameter " + c.Param("id")})
		return
	}

	if bookID <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "the book you are looking for does not exist"})
		return
	}

	var book Book

	errors := db.First(&book, bookID).GetErrors()

	for _, err := range errors {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "the book you are looking for does not exist"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}

	if book.ID <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "the book you are looking for does not exist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": book})
}

// FindBookByName get book message by book name
func FindBookByName(c *gin.Context) {
	db, err = gorm.Open(dbConn.DBEngine, dbConn.DBName)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}

	defer db.Close()

	bookName := c.Param("name")

	var book Book

	visible := 1
	errors := db.Where("book_name = ? and visibility = ?", bookName, visible).First(&book).GetErrors()

	for _, err := range errors {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "the book you are looking for does not exist"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}

	if book.ID <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "the book you are looking for does not exist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": book})
}

// GetAllBooks Get All Books
func GetAllBooks(c *gin.Context) {
	db, err = gorm.Open(database.DBEngine, database.DBName)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	var bookList []Book
	visible := 1
	db.Where("visibility = ?", visible).Find(&bookList)

	c.JSON(http.StatusOK, gin.H{
		"message": bookList,
	})
}

// BorrowBook borrow a book
func BorrowBook(c *gin.Context) {

	db, err = gorm.Open(database.DBEngine, database.DBName)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	var userRequest UserRequest
	err := c.BindJSON(&userRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post data error!"})
	}

	var book Book

	bookID := userRequest.BookID
	userID := userRequest.UserID

	fmt.Println(bookID)
	fmt.Println(userID)

	visible := 1
	errors := db.Where("id = ? and visibility = ?", bookID, visible).First(&book).GetErrors()

	for _, err := range errors {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "The book you are looking for does not exist!"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}

	var currentUser user.User
	userError := db.Where("id = ?", userID).First(&currentUser).GetErrors()

	for _, err := range userError {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "The user does not exist"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}

	if book.BookStatus == constant.BookBusy {
		c.JSON(http.StatusNotFound, gin.H{"error": "The book is already been borrowed"})
		return
	}

	book.BookStatus = constant.BookBusy
	book.LeftAmount--
	book.UpdatedTime = time.Now()

	db.Save(&book)

	var bookHistory history.BorrowHistory

	bookHistory.UserID = currentUser.ID
	bookHistory.UserName = currentUser.Username
	bookHistory.BookID = bookID
	bookHistory.BookName = book.BookName
	bookHistory.BorrowDate = time.Now()
	bookHistory.HistoryStatus = constant.BookFree
	bookHistory.CreatedTime = time.Now()
	bookHistory.UpdatedTime = time.Now()

	db.Save(&bookHistory)

	//check book status(bookId)
	//update book left status(bookId)
	//add a history record(bookId, userId)
	//return the book user borrowed

	c.JSON(http.StatusOK, gin.H{"message": book})
}

// ReturnBook return a book
func ReturnBook(c *gin.Context) {

	db, err = gorm.Open(database.DBEngine, database.DBName)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	var userRequest UserRequest
	err := c.BindJSON(&userRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post data error!"})
	}

	var book Book

	bookID := userRequest.BookID
	userID := userRequest.UserID

	errors := db.First(&book, bookID).GetErrors()

	for _, err := range errors {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "the book you are looking for does not exist"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}

	var bookHistory history.BorrowHistory
	bookErros := db.Where("user_id = ? and user_id = ? and history_status = ?", userID, bookID, constant.BookFree).First(&bookHistory).GetErrors()

	for _, err := range bookErros {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "No such book record"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}

	bookHistory.ReturnDate = time.Now()
	bookHistory.UpdatedTime = time.Now()
	bookHistory.HistoryStatus = constant.BookBusy

	db.Save(&bookHistory)

	book.BookStatus = constant.BookFree
	book.UpdatedTime = time.Now()
	book.LeftAmount++
	db.Save(&book)

	c.JSON(http.StatusOK, gin.H{"message": bookHistory})
}

// AddBook : add a new book
func AddBook(c *gin.Context) {
	db, err = gorm.Open(database.DBEngine, database.DBName)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	var bookRequest BookRequest
	err := c.BindJSON(&bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post data error!"})
		return
	}

	var book Book
	book.BookName = bookRequest.BookName
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

	db.Save(&book)

	c.JSON(http.StatusOK, gin.H{"message": book})
	return
}

// UpdateBookInfo : update book information
func UpdateBookInfo(c *gin.Context) {

	db, err = gorm.Open(database.DBEngine, database.DBName)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	var bookRequest BookRequest
	err := c.BindJSON(&bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post data error!"})
		return
	}

	if bookRequest.BookID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book id required"})
		return
	}

	var book Book
	bookErrors := db.Where("id = ?", bookRequest.BookID).First(&book).GetErrors()

	for _, err := range bookErrors {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "No such book record"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}

	if bookRequest.BookName != "" {
		book.BookName = bookRequest.BookName
	}
	if bookRequest.Description != "" {
		book.BookDescription = bookRequest.Description
	}
	if bookRequest.Author != "" {
		book.Author = bookRequest.Author
	}
	book.UpdatedTime = time.Now()

	db.Save(&book)

	c.JSON(http.StatusOK, gin.H{"message": book})
	return
}

// DeleteBook : logic delete, just make it invisible
func DeleteBook(c *gin.Context) {

	db, err = gorm.Open(dbConn.DBEngine, dbConn.DBName)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}

	defer db.Close()

	bookID := c.Param("bookID")

	var book Book

	visible := 1

	errors := db.Where("id = ? and visibility = ?", bookID, visible).First(&book).GetErrors()

	for _, err := range errors {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "No such book record"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}

	book.Visibility = 0
	db.Save(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
