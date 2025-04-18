package auth

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/kyimmQ/ielts-writing-golang/internal/modules/auth/dto"
	userDTO "github.com/kyimmQ/ielts-writing-golang/internal/modules/user/dto"
	errors "github.com/kyimmQ/ielts-writing-golang/pkg/error"
	"github.com/kyimmQ/ielts-writing-golang/pkg/response"
)

type AuthHandlerI interface {
	SignUp(ctx *gin.Context)
	SignIn(ctx *gin.Context)
}

type AuthHandler struct {
	authService AuthServiceI
}

func NewAuthHandler(authService AuthServiceI) AuthHandlerI {
	return &AuthHandler{
		authService: authService,
	}
}

// @Summary		User register account to system.
// @Description	Register account to system.
// @Tags			Policy
// @Accept			json
// @Produce		json
// @Param			CreateUserRequest	body		userDTO.CreateUserRequest	true	"CreateUserRequest JSON"
// @Success		200	{object}	response.SuccessResponse{data=string}
// @Failure	400 {object} response.ErrorResponse "{"errorKey": "ERR_INVALID_INPUT", "errorMessage": "invalid input"}"
// @Failure	500 {object} response.ErrorResponse "{"errorKey": "UserHashPasswordError", "errorMessage": "failed to hash password"}"
// @Failure	500 {object} response.ErrorResponse "{"errorKey": "UserAddToDBError", "errorMessage": "failed to add user to database"}"
// @Router			/auth/signup [post]
func (h *AuthHandler) SignUp(ctx *gin.Context) {
	var req userDTO.CreateUserRequest

	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		slog.Error("Failed to sign up user", "error", err)
		response.ResponseError(ctx, errors.ErrInvalidInput(err))
		return
	}

	if err := h.authService.SignUp(ctx, &req); err != nil {
		slog.Error("Failed to sign up user", "error", err)
		response.ResponseError(ctx, err)
		return
	}

	slog.Info("User signed up successfully", "email", req.Email)
	response.ReponseSuccess(ctx, "User created successfully", nil)
}

// @Summary		User sign in system.
// @Description	Sign user into system.
// @Tags			Policy
// @Accept			json
// @Produce		json
// @Param			SignInRequest	body		dto.SignInRequest	true	"SignInRequest JSON"
// @Success		200	{object}	response.SuccessResponse{data=dto.SignInResponse}
// @Failure	400 {object} response.ErrorResponse "{"errorKey": "ERR_INVALID_INPUT", "errorMessage": "invalid input"}"
// @Failure	401 {object} response.ErrorResponse "{"errorKey": "AuthInvalidPassword", "errorMessage": "invalid password"}"
// @Failure	404 {object} response.ErrorResponse "{"errorKey": "UserNotFound", "errorMessage": "username not found"}"
// @Failure	500 {object} response.ErrorResponse "{"errorKey": "ERR_GENERATE_TOKEN", "errorMessage": "string"}"
// @Router			/auth/signin [post]
func (h *AuthHandler) SignIn(ctx *gin.Context) {
	var req dto.SignInRequest

	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		slog.Error("Failed to parse request", "error", err)
		response.ResponseError(ctx, errors.ErrInvalidInput(err))
		return
	}

	token, err := h.authService.SignIn(ctx, &req)

	if err != nil {
		slog.Error("Failed to sign in user", "error", err)
		response.ResponseError(ctx, err)
		return
	}

	var responseData dto.SignInResponse
	responseData.Token = token

	slog.Info("User signed in successfully", "username", req.Username)
	response.ReponseSuccess(ctx, "User signed in successfully", responseData)

}
