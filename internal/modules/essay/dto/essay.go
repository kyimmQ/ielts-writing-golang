package dto

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kyimmQ/ielts-writing-golang/internal/entity"
	"github.com/kyimmQ/ielts-writing-golang/pkg/utils"
)

type CreateEssayRequest struct {
	PromptID  uuid.UUID `json:"promptId" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	Status    string    `json:"status" validate:"required,oneof=draft grading"`
	TimeTaken int       `json:"timeTaken,omitempty"`
}

func (req *CreateEssayRequest) ToEntity(ctx context.Context) *entity.Essay {
	newUUID, _ := utils.GenerateUUID()

	userUUID := ctx.Value("userId").(uuid.UUID)
	return &entity.Essay{
		ID:        newUUID,
		UserID:    userUUID,
		PromptID:  req.PromptID,
		Content:   req.Content,
		Status:    entity.EssayStatus(req.Status),
		TimeTaken: req.TimeTaken,
		UpdatedAt: time.Now(),
	}
}

type UpdateEssayRequest struct {
	ID        uuid.UUID `json:"id" bson:"_id"`
	Content   string    `json:"content" bson:"content"`
	TimeTaken int       `json:"timeTaken,omitempty" bson:"band,omitempty"`
}

type EssayWithPromptResponse struct {
	ID        uuid.UUID       `json:"id"`
	PromptID  uuid.UUID       `json:"promptId"`
	Prompt    string          `json:"prompt"`
	Content   string          `json:"content"`
	Status    string          `json:"status"`
	Band      entity.BandType `json:"band,omitempty"`
	TimeTaken int             `json:"timeTaken,omitempty"`
	UpdatedAt time.Time       `json:"updatedAt"`
}
