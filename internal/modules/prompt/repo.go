package prompt

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/kyimmQ/ielts-writing-golang/global"
	"github.com/kyimmQ/ielts-writing-golang/internal/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var (
	CollectionName = "prompts"
)

type PromptRepositoryI interface {
	CreatePrompt(ctx context.Context, prompt *entity.ExamPrompt) error
	GetRandomPrompt(ctx context.Context) (*entity.ExamPrompt, error)
	GetPromptByID(ctx context.Context, id uuid.UUID) (*entity.ExamPrompt, error)
}

type PromptRepository struct {
	db *mongo.Client
}

func NewPromptRepository(db *mongo.Client) PromptRepositoryI {
	return &PromptRepository{db: db}
}

func (r *PromptRepository) CreatePrompt(ctx context.Context, prompt *entity.ExamPrompt) error {
	collection := r.db.Database(global.Config.MongoDB.DatabaseName).Collection(CollectionName)
	_, err := collection.InsertOne(ctx, prompt)
	if err != nil {
		return fmt.Errorf("fail to create new prompt, error: %v", err)
	}
	return nil
}

func (r *PromptRepository) GetRandomPrompt(ctx context.Context) (*entity.ExamPrompt, error) {
	collection := r.db.Database(global.Config.MongoDB.DatabaseName).Collection(CollectionName)

	// Define the aggregation pipeline with the $sample stage
	sampleStage := bson.D{{Key: "$sample", Value: bson.D{{Key: "size", Value: 1}}}}
	// Count total prompts
	cursor, err := collection.Aggregate(ctx, mongo.Pipeline{sampleStage})
	if err != nil {
		return nil, fmt.Errorf("fail to get random prompt, error: %v", err)
	}
	defer cursor.Close(ctx)

	var result []entity.ExamPrompt
	if err = cursor.All(ctx, &result); err != nil {
		return nil, fmt.Errorf("error parsing exam prompt, error: %v", err)
	}

	return &(result[0]), nil

}

func (r *PromptRepository) GetPromptByID(ctx context.Context, id uuid.UUID) (*entity.ExamPrompt, error) {
	collection := r.db.Database(global.Config.MongoDB.DatabaseName).Collection(CollectionName)

	var prompt entity.ExamPrompt
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&prompt)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("prompt with ID %s not found", id)
		}
		return nil, fmt.Errorf("failed to get prompt by ID %s, error: %v", id, err)
	}
	return &prompt, nil
}
