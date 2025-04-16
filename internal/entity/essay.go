package entity

import (
	"time"

	"github.com/google/uuid"
)

type BandType float32

const (
	Band_0  BandType = 0.0
	Band_05 BandType = 0.5
	Band_1  BandType = 1.0
	Band_15 BandType = 1.5
	Band_2  BandType = 2.0
	Band_25 BandType = 2.5
	Band_3  BandType = 3.0
	Band_35 BandType = 3.5
	Band_4  BandType = 4.0
	Band_45 BandType = 4.5
	Band_5  BandType = 5.0
	Band_55 BandType = 5.5
	Band_6  BandType = 6.0
	Band_65 BandType = 6.5
	Band_7  BandType = 7.0
	Band_75 BandType = 7.5
	Band_8  BandType = 8.0
	Band_85 BandType = 8.5
	Band_9  BandType = 9.0
)

type EssayStatus string

const (
	StatusDraft   EssayStatus = "draft"
	StatusGrading EssayStatus = "grading"
	StatusGraded  EssayStatus = "graded"
)

type Essay struct {
	ID        uuid.UUID   `json:"id" bson:"_id"`
	UserID    uuid.UUID   `json:"userId" bson:"userId"`
	PromptID  uuid.UUID   `json:"promptId" bson:"promptId"`
	Content   string      `json:"content" bson:"content"`
	Status    EssayStatus `json:"status" bson:"status"`
	Band      BandType    `json:"band,omitempty" bson:"band,omitempty"`
	TimeTaken int         `json:"timeTaken,omitempty" bson:"band,omitempty"`
	UpdatedAt time.Time   `json:"updatedAt" bson:"updatedAt"`
}
