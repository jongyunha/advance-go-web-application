package dto

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
