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

type DonationRepositoryImpl struct {
	collection *mongo.Collection
	config     *config.Config
}

func NewDonationRepositoryImpl(db *mongo.Database) infrastructureInterface.DonationRepository {
	return &DonationRepositoryImpl{
		collection: db.Collection("donations"),
		config:     config.Load(),
	}
}

func (r *DonationRepositoryImpl) CreateDonation(ctx context.Context, donation entity.Donation) (primitive.ObjectID, error) {
	result, err := r.collection.InsertOne(ctx, donation)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *DonationRepositoryImpl) GetDonationByID(ctx context.Context, id primitive.ObjectID) (*entity.Donation, error) {
	var donation entity.Donation
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&donation)
	if err != nil {
		return nil, err
	}
	return &donation, nil
}

func (r *DonationRepositoryImpl) GetDonationsByCampaign(ctx context.Context, campaignID primitive.ObjectID) ([]entity.Donation, error) {
	var donations []entity.Donation
	cursor, err := r.collection.Find(ctx, bson.M{"campaign": campaignID})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &donations); err != nil {
		return nil, err
	}
	return donations, nil
}

func (r *DonationRepositoryImpl) GetDonationsByDonor(ctx context.Context, donorID primitive.ObjectID) ([]entity.Donation, error) {
	var donations []entity.Donation
	cursor, err := r.collection.Find(ctx, bson.M{"donor": donorID})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &donations); err != nil {
		return nil, err
	}
	return donations, nil
}

func (r *DonationRepositoryImpl) DeleteDonation(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
