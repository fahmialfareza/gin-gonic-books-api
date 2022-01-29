package dto

// LoginDTO is a struct for login
type LoginDTO struct {
	Email    string `json:"email" binding:"required" form:"email" validate:"email"`
	Password string `json:"password" form:"password" validate:"min=8" binding:"required"`
}
