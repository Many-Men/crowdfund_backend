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

type CampaignRepositoryImpl struct {
	collection *mongo.Collection
	config     *config.Config
}

func NewCampaignRepositoryImpl(db *mongo.Database) infrastructureInterface.CampaignRepository {
	return &CampaignRepositoryImpl{
		collection: db.Collection("campaigns"),
		config:     config.Load(),
	}
}

func (r *CampaignRepositoryImpl) CreateCampaign(ctx context.Context, campaign entity.Campaign) (primitive.ObjectID, error) {
	result, err := r.collection.InsertOne(ctx, campaign)
	if err != nil {
		return primitive.NilObjectID, &_errors.InternalServerError{Message: "Failed to create campaign"}
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *CampaignRepositoryImpl) GetCampaignByID(ctx context.Context, id primitive.ObjectID) (*entity.Campaign, error) {
	var campaign entity.Campaign
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&campaign)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, &_errors.NotFoundError{Message: "Campaign not found"}
		}
		return nil, &_errors.InternalServerError{Message: "Failed to retrieve campaign"}
	}
	return &campaign, nil
}

func (r *CampaignRepositoryImpl) GetAllCampaigns(ctx context.Context) ([]entity.Campaign, error) {
	var campaigns []entity.Campaign
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, &_errors.InternalServerError{Message: "Failed to retrieve campaigns"}
	}
	if err = cursor.All(ctx, &campaigns); err != nil {
		return nil, &_errors.InternalServerError{Message: "Failed to process campaigns data"}
	}
	return campaigns, nil
}

func (r *CampaignRepositoryImpl) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User
	err := r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, &_errors.NotFoundError{Message: "User not found"}
		}
		return nil, &_errors.InternalServerError{Message: "Failed to retrieve user"}
	}
	return &user, nil
}

func (r *CampaignRepositoryImpl) UpdateCampaignAmount(ctx context.Context, id primitive.ObjectID, amount float64) error {
	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"current_amount": amount}})
	if err != nil {
		return &_errors.InternalServerError{Message: "Failed to update campaign amount"}
	}
	if result.MatchedCount == 0 {
		return &_errors.NotFoundError{Message: "Campaign not found"}
	}
	return nil
}

func (r *CampaignRepositoryImpl) DeleteCampaign(ctx context.Context, id primitive.ObjectID) error {
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return &_errors.InternalServerError{Message: "Failed to delete campaign"}
	}
	if result.DeletedCount == 0 {
		return &_errors.NotFoundError{Message: "Campaign not found"}
	}
	return nil
}
