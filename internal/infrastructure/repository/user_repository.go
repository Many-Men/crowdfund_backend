package repository

import (
	"context"
	"errors"
	"github.com/Many-Men/crowdfund_backend/config"
	_errors "github.com/Many-Men/crowdfund_backend/errors"
	"github.com/Many-Men/crowdfund_backend/internal/infrastructure/entity"
	infrastructureInterface "github.com/Many-Men/crowdfund_backend/internal/service/interface"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImpl struct {
	collection *mongo.Collection
	config     *config.Config
}

func NewUserRepositoryImpl(db *mongo.Database) infrastructureInterface.UserRepository {
	return &UserRepositoryImpl{
		collection: db.Collection("user"),
		config:     config.Load(),
	}
}

func (r *UserRepositoryImpl) CreateUser(ctx context.Context, user entity.User) (primitive.ObjectID, error) {
	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, &_errors.InternalServerError{Message: "Failed to create user"}
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *UserRepositoryImpl) GetUserByID(ctx context.Context, id primitive.ObjectID) (*entity.User, error) {
	var user entity.User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, &_errors.NotFoundError{Message: "User not found"}
		}
		return nil, &_errors.InternalServerError{Message: "Failed to retrieve user"}
	}
	return &user, nil
}

func (r *UserRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, &_errors.NotFoundError{Message: "User not found"}
		}
		return nil, &_errors.InternalServerError{Message: "Failed to retrieve user"}
	}
	return &user, nil
}

func (r *UserRepositoryImpl) UpdateUserBalance(ctx context.Context, id primitive.ObjectID, balance float64) error {
	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"balance": balance}})
	if err != nil {
		return &_errors.InternalServerError{Message: "Failed to update user balance"}
	}
	if result.MatchedCount == 0 {
		return &_errors.NotFoundError{Message: "User not found"}
	}
	return nil
}

func (r *UserRepositoryImpl) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return &_errors.InternalServerError{Message: "Failed to delete user"}
	}
	if result.DeletedCount == 0 {
		return &_errors.NotFoundError{Message: "User not found"}
	}
	return nil
}

func (r *UserRepositoryImpl) ListUsers(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, &_errors.InternalServerError{Message: "Failed to retrieve users"}
	}
	if err = cursor.All(ctx, &users); err != nil {
		return nil, &_errors.InternalServerError{Message: "Failed to process users data"}
	}
	return users, nil
}
