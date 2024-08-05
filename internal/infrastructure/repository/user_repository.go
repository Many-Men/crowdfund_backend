package repository

import (
	"context"
	"github.com/Many-Men/crowdfund_backend/config"
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
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *UserRepositoryImpl) GetUserByID(ctx context.Context, id primitive.ObjectID) (*entity.User, error) {
	var user entity.User
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) UpdateUserBalance(ctx context.Context, id primitive.ObjectID, balance float64) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"balance": balance}})
	return err
}

func (r *UserRepositoryImpl) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *UserRepositoryImpl) ListUsers(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}
