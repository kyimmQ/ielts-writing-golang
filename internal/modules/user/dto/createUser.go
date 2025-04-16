package dto

import "github.com/kyimmQ/ielts-writing-golang/internal/entity"

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

func (req *CreateUserRequest) ToEntity() *entity.User {
	return &entity.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
}
