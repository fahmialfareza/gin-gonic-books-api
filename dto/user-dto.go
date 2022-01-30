package dto

// UserUpdateDTO is a data transfer object for updating a user.
type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id"`
	Name     string `json:"name" binding:"required" form:"name"`
	Email    string `json:"email" binding:"required,email" form:"email""`
	Password string `json:"password,omitempty" form:"password,omitempty"`
}
