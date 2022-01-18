package dto

type LoginDTO struct {
	Email    string `json:"email" form:"email" validate:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required" validate:"min-6"`
}
