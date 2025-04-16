package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errors "github.com/kyimmQ/ielts-writing-golang/pkg/error"
)

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ReponseSuccess(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, SuccessResponse{
		Message: message,
		Data:    data,
	})
}

type ErrorResponse struct {
	ErrorKey     string `json:"errorKey"`
	ErrorMessage string `json:"errorMessage"`
}

func ResponseError(c *gin.Context, err error) {
	if domainErr, ok := err.(*errors.DomainError); ok {
		c.JSON(domainErr.StatusCode, ErrorResponse{
			ErrorKey:     domainErr.ErrorKey,
			ErrorMessage: domainErr.Message,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, ErrorResponse{
		ErrorKey:     "InternalError",
		ErrorMessage: "An internal error occurred",
	})
}
