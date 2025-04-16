package dto

import "github.com/kyimmQ/ielts-writing-golang/internal/entity"

type SignUpRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
}

func (req *SignUpRequest) ToEntity() *entity.User {
	return &entity.User{
		Username: req.Username,
		Email:    req.Email,
		Name:     req.Name,
	}
}
