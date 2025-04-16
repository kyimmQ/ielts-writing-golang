package entity

import (
	"github.com/google/uuid"
)

type ExamPrompt struct {
	ID     uuid.UUID `json:"id" bson:"_id"`
	Prompt string    `json:"prompt" bson:"prompt"`
}
