package middleware

import (
	"log/slog"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kyimmQ/ielts-writing-golang/global"
	errors "github.com/kyimmQ/ielts-writing-golang/pkg/error"
	"github.com/kyimmQ/ielts-writing-golang/pkg/jwt"
	"github.com/kyimmQ/ielts-writing-golang/pkg/response"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			slog.Error("Authorization header is missing")
			response.ResponseError(c, errors.ErrUnauthorized())
			c.Abort() // Stop further processing
			return
		}

		// Check Bearer prefix
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			response.ResponseError(c, errors.ErrUnauthorized())
			c.Abort() // Stop further processing
			return
		}

		// Validate token
		token, err := jwt.VerifyToken(tokenString, global.Config.JWT.SecretKey)
		if err != nil {
			response.ResponseError(c, errors.ErrUnauthorized())
			c.Abort() // Stop further processing
			return
		}

		// Set user information in the context
		c.Set("userId", token.UserID.String())
		// Proceed to the next middleware or handler
		c.Next()
	}
}
