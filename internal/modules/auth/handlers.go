package auth

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/kyimmQ/ielts-writing-golang/internal/modules/auth/dto"
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
// @Param			SignUpRequest	body		dto.SignUpRequest	true	"SignUpRequest JSON"
// @Success		200	{object}	response.SuccessResponse{data=string}
// @Failure	400 {object} response.ErrorResponse
// @Router			/auth/signup [post]
func (h *AuthHandler) SignUp(ctx *gin.Context) {
	var req dto.SignUpRequest

	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		slog.Error("Failed to sign up user", "error", err)
		response.ResponseError(ctx, errors.ErrInvalidInput(err))
		return
	}

	if err := h.authService.SignUp(&req); err != nil {
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
// @Success		200	{object}	response.SuccessResponse{data=string}
// @Failure	400 {object} response.ErrorResponse
// @Router			/auth/signin [post]
func (h *AuthHandler) SignIn(ctx *gin.Context) {
	var req dto.SignInRequest

	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		slog.Error("Failed to parse request", "error", err)
		response.ResponseError(ctx, errors.ErrInvalidInput(err))
		return
	}

	token, err := h.authService.SignIn(&req)
	if err != nil {
		slog.Error("Failed to sign in user", "error", err)
		response.ResponseError(ctx, err)
		return
	}

	slog.Info("User signed in successfully", "username", req.Username)
	response.ReponseSuccess(ctx, "User signed in successfully", token)

}
