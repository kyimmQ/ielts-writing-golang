package auth

import (
	"context"
	"fmt"

	"github.com/kyimmQ/ielts-writing-golang/internal/modules/auth/dto"
	"github.com/kyimmQ/ielts-writing-golang/internal/modules/auth/helper"
	"github.com/kyimmQ/ielts-writing-golang/internal/modules/user"
	"github.com/kyimmQ/ielts-writing-golang/pkg/hash"
)

type AuthServiceI interface {
	SignUp(ctx context.Context, req *dto.SignUpRequest) error
	SignIn(ctx context.Context, req *dto.SignInRequest) (string, error)
}

type AuthService struct {
	usrRepo user.UserRepositoryI
}

func NewAuthService(usrRepo user.UserRepositoryI) AuthServiceI {
	return &AuthService{
		usrRepo: usrRepo,
	}
}

func (s *AuthService) SignUp(ctx context.Context, req *dto.SignUpRequest) error {
	// Convert the request to a user entity
	userEntity := req.ToEntity()

	// Hash the password
	hashPassword, err := hash.Generate(req.Password)
	if err != nil {
		return err
	}

	// Set the hashed password in the user entity
	userEntity.Password = hashPassword
	// Call the user repository to create the user
	err = s.usrRepo.CreateUser(userEntity)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) SignIn(ctx context.Context, req *dto.SignInRequest) (string, error) {
	// Find the user by username
	userEntity, err := s.usrRepo.FindUserByUsername(req.Username)
	if err != nil {
		return "", err
	}

	// Check if the password is correct
	if ok := hash.Validate(userEntity.Password, req.Password); !ok {
		return "", fmt.Errorf("invalid password")
	}

	token, err := helper.GenerateAccessToken(userEntity.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
