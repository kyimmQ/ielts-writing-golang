package dto

import (
	"github.com/kyimmQ/ielts-writing-golang/internal/entity"
	"github.com/kyimmQ/ielts-writing-golang/pkg/utils"
)

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
}

func (req *CreateUserRequest) ToEntity() *entity.User {
	newUUID, _ := utils.GenerateUUID()
	return &entity.User{
		ID:       newUUID,
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
}
