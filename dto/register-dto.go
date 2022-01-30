package dto

// RegisterDTO is a data transfer object for registering a new user.
type RegisterDTO struct {
	Name     string `json:"name" binding:"required,min=1" form:"name"`
	Email    string `json:"email" binding:"required" form:"email" validate:"email"`
	Password string `json:"password" form:"password" binding:"required,min=8"`
}
