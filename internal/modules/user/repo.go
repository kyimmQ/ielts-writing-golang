package user

import (
	"context"
	"fmt"

	"github.com/kyimmQ/ielts-writing-golang/internal/entity"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var (
	dbName         = "ielts-writing"
	collectionName = "users"
)

type UserRepositoryI interface {
	CreateUser(user *entity.User) error
	FindUserByUsername(username string) (*entity.User, error)
}

type UserRepository struct {
	db *mongo.Client
}

func NewUserRepository(db *mongo.Client) UserRepositoryI {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user *entity.User) error {
	// Insert the user into the database
	collection := r.db.Database(dbName).Collection(collectionName)
	_, err := collection.InsertOne(context.TODO(), &user)
	if err != nil {
		return fmt.Errorf("failed to create user: %s, error: %v", user.Username, err)
	}
	return nil
}

func (r *UserRepository) FindUserByUsername(username string) (*entity.User, error) {
	// Find the user by username
	collection := r.db.Database(dbName).Collection(collectionName)
	filter := bson.D{{Key: "username", Value: username}}
	var user entity.User
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to find user by username: %s, error: %v", username, err)
	}
	return &user, nil
}
