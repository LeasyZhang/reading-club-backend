package dto

// BookRequest request body of book api
type BookRequest struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Author      string  `json:"author"`
	ISBN        string  `json:"ISBN"`
	DoubanURL   string  `json:"doubanUrl"`
	ImageURL    string  `json:"imageUrl"`
	Price       float32 `json:"price"`
	Press       string  `json:"press"`
	Description string  `json:"description"`
}
