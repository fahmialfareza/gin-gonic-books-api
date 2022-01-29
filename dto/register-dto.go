package dto

// RegisterDTO is a data transfer object for registering a new user.
type RegisterDTO struct {
	Name     string `json:"name" binding:"required" form:"name" validate:"min=1"`
	Email    string `json:"email" binding:"required" form:"email" validate:"email"`
	Password string `json:"password" form:"password" validate:"min=8" binding:"required"`
}
