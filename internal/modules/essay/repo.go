package essay

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/kyimmQ/ielts-writing-golang/global"
	"github.com/kyimmQ/ielts-writing-golang/internal/entity"
	"github.com/kyimmQ/ielts-writing-golang/internal/modules/essay/dto"
	errors "github.com/kyimmQ/ielts-writing-golang/pkg/error"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type EssayRepositoryI interface {
	CreateEssay(ctx context.Context, e *entity.Essay) error
	GetUserEssays(ctx context.Context, userID uuid.UUID, statusFilter []entity.EssayStatus) ([]entity.Essay, error)
	UpdateEssayDraft(ctx context.Context, e *dto.UpdateEssayRequest) error
}

type EssayRepository struct {
	db *mongo.Client
}

func NewEssayRepository(db *mongo.Client) EssayRepositoryI {
	return &EssayRepository{db: db}
}

const (
	collectionName = "essays"
)

func (r *EssayRepository) CreateEssay(ctx context.Context, e *entity.Essay) error {
	coll := r.db.Database(global.Config.MongoDB.DatabaseName).Collection(collectionName)
	_, err := coll.InsertOne(ctx, e)
	if err != nil {
		return errors.NewDomainError(http.StatusInternalServerError, err, "failed to create essay", "CreateEssayError")
	}
	return nil
}

func (r *EssayRepository) GetUserEssays(ctx context.Context, userID uuid.UUID, statuses []entity.EssayStatus) ([]entity.Essay, error) {
	essayCollection := r.db.Database(global.Config.MongoDB.DatabaseName).Collection(collectionName)

	filter := bson.M{
		"userId": userID,
		"status": bson.M{"$in": statuses},
	}

	cursor, err := essayCollection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to get essay(s), error: %v", err)
	}

	var results []entity.Essay

	if err := cursor.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("failed to parse essay(s), error: %v", err)
	}

	return results, nil
}

func (r *EssayRepository) UpdateEssayDraft(ctx context.Context, e *dto.UpdateEssayRequest) error {
	filter := bson.M{"_id": e.ID, "status": entity.StatusDraft}
	update := bson.M{
		"$set": bson.M{
			"content":   e.Content,
			"timeTaken": e.TimeTaken,
			"updatedAt": time.Now(),
		},
	}
	_, err := r.db.Database(global.Config.MongoDB.DatabaseName).Collection(collectionName).UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update essay, error: %v", err)
	}
	return nil
}
