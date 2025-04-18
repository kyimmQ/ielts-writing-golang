package auth

import (
	"context"
	"net/http"

	"github.com/kyimmQ/ielts-writing-golang/internal/modules/auth/dto"
	"github.com/kyimmQ/ielts-writing-golang/internal/modules/auth/helper"
	"github.com/kyimmQ/ielts-writing-golang/internal/modules/user"
	userDTO "github.com/kyimmQ/ielts-writing-golang/internal/modules/user/dto"
	errors "github.com/kyimmQ/ielts-writing-golang/pkg/error"
	"github.com/kyimmQ/ielts-writing-golang/pkg/hash"
)

type AuthServiceI interface {
	SignUp(ctx context.Context, req *userDTO.CreateUserRequest) error
	SignIn(ctx context.Context, req *dto.SignInRequest) (string, error)
}

type AuthService struct {
	usrService user.UserServiceI
}

func NewAuthService(usrService user.UserServiceI) AuthServiceI {
	return &AuthService{
		usrService: usrService,
	}
}

func (s *AuthService) SignUp(ctx context.Context, req *userDTO.CreateUserRequest) error {
	// Call the user repository to create the user
	err := s.usrService.CreateUser(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) SignIn(ctx context.Context, req *dto.SignInRequest) (string, error) {
	// Find the user by username
	var findUserReq userDTO.GetUserByUsernameRequest
	findUserReq.Username = req.Username
	userEntity, err := s.usrService.GetUserByUsername(ctx, &findUserReq)
	if err != nil {
		return "", err
	}

	// Check if the password is correct
	if ok := hash.Validate(userEntity.Password, req.Password); !ok {
		return "", errors.NewDomainError(http.StatusUnauthorized, nil, "invalid password", "AuthInvalidPassword")
	}

	token, err := helper.GenerateAccessToken(userEntity.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
