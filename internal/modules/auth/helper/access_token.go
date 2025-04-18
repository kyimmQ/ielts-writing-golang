package helper

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/kyimmQ/ielts-writing-golang/global"
	errors "github.com/kyimmQ/ielts-writing-golang/pkg/error"
	"github.com/kyimmQ/ielts-writing-golang/pkg/jwt"
)

func GenerateAccessToken(id uuid.UUID) (string, error) {
	accessSecretKey := global.Config.JWT.SecretKey

	accessExpiry, err := strconv.Atoi(global.Config.JWT.Expiry)
	if err != nil {
		return "", errors.NewDomainError(http.StatusInternalServerError, nil, "cannot parse expiry token", "ERR_GENERATE_TOKEN")
	}

	accessToken, err := jwt.GenerateToken(accessSecretKey, accessExpiry, id)
	if err != nil {

		return "", errors.NewDomainError(http.StatusInternalServerError, nil, "cannot generate access token", "ERR_GENERATE_TOKEN")
	}

	return accessToken, nil
}
