package dto

// LoginDTO is a struct for login
type LoginDTO struct {
	Email    string `json:"email" binding:"required,email" form:"email""`
	Password string `json:"password" form:"password" binding:"required,min=8"`
}
