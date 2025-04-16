package user

import (
	"github.com/kyimmQ/ielts-writing-golang/internal/modules/user/dto"
)

type UserServiceI interface {
	CreateUser(req *dto.CreateUserRequest) error
}

type UserSevice struct {
	userRepo UserRepositoryI
}

func NewUserService(userRepo UserRepository) UserServiceI {
	return &UserSevice{
		userRepo: &userRepo,
	}
}

func (s *UserSevice) CreateUser(req *dto.CreateUserRequest) error {
	user := req.ToEntity()
	err := s.userRepo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}
