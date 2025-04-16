package dto

import (
	"github.com/google/uuid"
	"github.com/kyimmQ/ielts-writing-golang/internal/entity"
	"github.com/kyimmQ/ielts-writing-golang/pkg/utils"
)

type CreatePromptRequest struct {
	Prompt string `json:"prompt" validate:"required"`
}

func (req *CreatePromptRequest) ToEntity() *entity.ExamPrompt {
	newUUID, _ := utils.GenerateUUID()
	return &entity.ExamPrompt{
		ID:     newUUID,
		Prompt: req.Prompt,
	}
}

type PromptResponse struct {
	ID     uuid.UUID `json:"id"`
	Prompt string    `json:"prompt"`
}
