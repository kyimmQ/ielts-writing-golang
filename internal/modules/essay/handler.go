package essay

import (
	"github.com/gin-gonic/gin"
	"github.com/kyimmQ/ielts-writing-golang/internal/modules/essay/dto"
	"github.com/kyimmQ/ielts-writing-golang/pkg/response"
)

type EssayHandlerI interface {
	SubmitEssay(ctx *gin.Context)
	SaveDraft(ctx *gin.Context)
	GetUserHistory(ctx *gin.Context)
	GetUserDrafts(ctx *gin.Context)
}

type EssayHandler struct {
	service EssayServiceI
}

func NewEssayHandler(service EssayServiceI) EssayHandlerI {
	return &EssayHandler{service: service}
}

// @Summary		Submit an essay for drafting or grading.
// @Description	User submits an essay for drafting or grading.
// @Tags			Essay
// @Security     Bearer
// @Accept			json
// @Produce		json
// @Param			CreateEssayRequest	body	dto.CreateEssayRequest	true	"Essay submission payload"
// @Success		200	{object}	response.SuccessResponse
// @Failure		400	{object}	response.ErrorResponse
// @Router			/essay/submit [post]
func (h *EssayHandler) SubmitEssay(ctx *gin.Context) {
	var req dto.CreateEssayRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ResponseError(ctx, err)
		return
	}

	err := h.service.SubmitEssay(ctx, &req)
	if err != nil {
		response.ResponseError(ctx, err)
		return
	}
	response.ReponseSuccess(ctx, "Essay submitted successfully", nil)
}

// @Summary		Save an essay draft
// @Description	User saves or updates a draft essay.
// @Tags			Essay
// @Security     Bearer
// @Accept			json
// @Produce		json
// @Param			UpdateEssayRequest	body	dto.UpdateEssayRequest	true	"Essay draft payload"
// @Success		200	{object}	response.SuccessResponse
// @Failure		400	{object}	response.ErrorResponse
// @Router			/essay/draft [put]
func (h *EssayHandler) SaveDraft(ctx *gin.Context) {
	var req dto.UpdateEssayRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ResponseError(ctx, err)
		return
	}

	err := h.service.SaveDraft(ctx, &req)
	if err != nil {
		response.ResponseError(ctx, err)
		return
	}
	response.ReponseSuccess(ctx, "Essay draft saved successfully", nil)
}

// @Summary		Get submitted essay history
// @Description	Get essays that have been submitted (grading or graded).
// @Tags			Essay
// @Security     Bearer
// @Produce		json
// @Success		200	{object}	response.SuccessResponse{data=[]dto.EssayWithPromptResponse}
// @Failure		400	{object}	response.ErrorResponse
// @Router			/essay/history [get]
func (h *EssayHandler) GetUserHistory(ctx *gin.Context) {
	essays, err := h.service.GetUserHistory(ctx)
	if err != nil {
		response.ResponseError(ctx, err)
		return
	}
	response.ReponseSuccess(ctx, "User essay history fetched", essays)
}

// @Summary		Get essay drafts
// @Description	Get all draft essays by the current user.
// @Tags			Essay
// @Security     Bearer
// @Produce		json
// @Success		200	{object}	response.SuccessResponse{data=[]dto.EssayWithPromptResponse}
// @Failure		400	{object}	response.ErrorResponse
// @Router			/essay/drafts [get]
func (h *EssayHandler) GetUserDrafts(ctx *gin.Context) {
	essays, err := h.service.GetUserDrafts(ctx)
	if err != nil {
		response.ResponseError(ctx, err)
		return
	}
	response.ReponseSuccess(ctx, "User draft essays fetched", essays)
}
