package user

import (
	"context"

	"github.com/kyimmQ/ielts-writing-golang/internal/modules/user/dto"
)

type UserServiceI interface {
	CreateUser(ctx context.Context, req *dto.CreateUserRequest) error
}

type UserSevice struct {
	userRepo UserRepositoryI
}

func NewUserService(userRepo UserRepositoryI) UserServiceI {
	return &UserSevice{
		userRepo: userRepo,
	}
}

func (s *UserSevice) CreateUser(ctx context.Context, req *dto.CreateUserRequest) error {
	user := req.ToEntity()
	err := s.userRepo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}
