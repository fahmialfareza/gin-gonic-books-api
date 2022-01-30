package dto

// UserUpdateDTO is a data transfer object for updating a user.
type UserUpdateDTO struct {
	ID       uint64 `json:"id" binding:"required" form:"id"`
	Name     string `json:"name" binding:"required" form:"name"`
	Email    string `json:"email" binding:"required" form:"email" validate:"email"`
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min=8"`
}

// UserCreateDTO is a data transfer object for creating a new user.
// type UserCreateDTO struct {
// 	Name     string `json:"name" binding:"required" form:"name"`
// 	Email    string `json:"email" binding:"required" form:"email" validate:"email"`
// 	Password string `json:"password" form:"password" validate:"min=8" binding:"required"`
// }
