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
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *CampaignRepositoryImpl) GetCampaignByID(ctx context.Context, id primitive.ObjectID) (*entity.Campaign, error) {
	var campaign entity.Campaign
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&campaign)
	if err != nil {
		return nil, err
	}
	return &campaign, nil
}

func (r *CampaignRepositoryImpl) GetAllCampaigns(ctx context.Context) ([]entity.Campaign, error) {
	var campaigns []entity.Campaign
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &campaigns); err != nil {
		return nil, err
	}
	return campaigns, nil
}

func (r *CampaignRepositoryImpl) UpdateCampaignAmount(ctx context.Context, id primitive.ObjectID, amount float64) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"current_amount": amount}})
	return err
}

func (r *CampaignRepositoryImpl) DeleteCampaign(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
