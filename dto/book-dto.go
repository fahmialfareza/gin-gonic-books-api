package dto

// BookUpdateDTO is a data transfer object for updating a book.
type BookUpdateDTO struct {
	ID          uint64 `json:"id" binding:"required" form:"id"`
	Title       string `json:"title" binding:"required" form:"title"`
	Description string `json:"description" binding:"required" form:"description"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}

// BookCreateDTO is a data transfer object for creating a book.
type BookCreateDTO struct {
	Title       string `json:"title" binding:"required" form:"title"`
	Description string `json:"description" binding:"required" form:"description"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}
