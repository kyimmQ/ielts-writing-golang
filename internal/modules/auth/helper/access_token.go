package helper

import (
	"errors"
	"strconv"

	"github.com/google/uuid"
	"github.com/kyimmQ/ielts-writing-golang/global"
	"github.com/kyimmQ/ielts-writing-golang/pkg/jwt"
)

func GenerateAccessToken(id uuid.UUID) (string, error) {
	accessSecretKey := global.Config.JWT.SecretKey

	accessExpiry, err := strconv.Atoi(global.Config.JWT.Expiry)
	if err != nil {
		return "", err
	}

	accessToken, err := jwt.GenerateToken(accessSecretKey, accessExpiry, id)
	if err != nil {

		return "", errors.New("failed to generate access token")
	}

	return accessToken, nil
}
