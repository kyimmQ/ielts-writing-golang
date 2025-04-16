package user

import (
	"context"
	"errors"

	"github.com/kyimmQ/ielts-writing-golang/internal/entity"
	"github.com/kyimmQ/ielts-writing-golang/internal/modules/user/dto"
	"github.com/kyimmQ/ielts-writing-golang/pkg/hash"
)

type UserServiceI interface {
	CreateUser(ctx context.Context, req *dto.CreateUserRequest) error
	GetUserByUsername(ctx context.Context, req *dto.GetUserByUsernameRequest) (*entity.User, error)
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
	// Convert the request to a user entity
	userEntity := req.ToEntity()

	// Hash the password
	hashPassword, err := hash.Generate(req.Password)
	if err != nil {
		return errors.New("error hasing password")
	}

	// Set the hashed password in the user entity
	userEntity.Password = hashPassword
	err = s.userRepo.CreateUser(ctx, userEntity)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserSevice) GetUserByUsername(ctx context.Context, req *dto.GetUserByUsernameRequest) (*entity.User, error) {
	return s.userRepo.GetUserByUsername(ctx, req.Username)
}
