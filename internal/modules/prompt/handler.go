package prompt

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/kyimmQ/ielts-writing-golang/internal/modules/prompt/dto"
	errors "github.com/kyimmQ/ielts-writing-golang/pkg/error"
	"github.com/kyimmQ/ielts-writing-golang/pkg/response"
)

type PromptHandlerI interface {
	CreatePrompt(ctx *gin.Context)
	GetRandomPrompt(ctx *gin.Context)
}

type PromptHandler struct {
	promptService PromptServiceI
}

func NewPromptHandler(promptService PromptServiceI) PromptHandlerI {
	return &PromptHandler{
		promptService: promptService,
	}
}

// CreatePrompt godoc
// @Summary		Create a new prompt
// @Description	Add a new IELTS writing prompt to the system.
// @Tags			Prompt
// @Security     Bearer
// @Accept			json
// @Produce		json
// @Param			CreatePromptRequest	body		dto.CreatePromptRequest	true	"Prompt data"
// @Success		200	{object}	response.SuccessResponse{data=string}
// @Failure		400	{object}	response.ErrorResponse
// @Router			/prompts [post]
func (h *PromptHandler) CreatePrompt(ctx *gin.Context) {
	var req dto.CreatePromptRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		slog.Error("Invalid prompt creation request", "error", err)
		response.ResponseError(ctx, errors.ErrInvalidInput(err))
		return
	}

	if err := h.promptService.CreatePrompt(ctx.Request.Context(), &req); err != nil {
		slog.Error("Failed to create prompt", "error", err)
		response.ResponseError(ctx, err)
		return
	}

	response.ReponseSuccess(ctx, "Prompt created successfully", nil)
}

// GetRandomPrompt godoc
// @Summary		Get a random prompt
// @Description	Retrieve a random IELTS writing prompt from the system.
// @Tags			Prompt
// @Security     Bearer
// @Produce		json
// @Success		200	{object}	response.SuccessResponse{data=dto.PromptResponse}
// @Failure		400	{object}	response.ErrorResponse
// @Router			/prompts/random [get]
func (h *PromptHandler) GetRandomPrompt(ctx *gin.Context) {
	prompt, err := h.promptService.GetRandomPrompt(ctx.Request.Context())
	if err != nil {
		slog.Error("Failed to get random prompt", "error", err)
		response.ResponseError(ctx, err)
		return
	}

	response.ReponseSuccess(ctx, "Prompt retrieved successfully", prompt)
}
