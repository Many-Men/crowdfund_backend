package _interface

import (
	"github.com/Many-Men/crowdfund_backend/internal/delivery/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserService interface {
	CreateUser(username string) error
	//GetUserByID(id primitive.ObjectID) (*model.UserResponse, error)
	//GetUserByEmail(email string) (*model.UserResponse, error)
	//UpdateUserBalance(id primitive.ObjectID, balance float64) error
	//DeleteUser(id primitive.ObjectID) error
	//ListUsers() ([]model.UserResponse, error)
}

type DonationService interface {
	//CreateDonation(donation model.DonationResponse) (primitive.ObjectID, error)
	//GetDonationByID(id primitive.ObjectID) (*model.DonationResponse, error)
	//GetDonationsByCampaign(campaignID primitive.ObjectID) ([]model.DonationResponse, error)
	//GetDonationsByDonor(donorID primitive.ObjectID) ([]model.DonationResponse, error)
	//DeleteDonation(id primitive.ObjectID) error
}

type CampaignService interface {
	CreateCampaign(title, description, username string, goal float64) error
	//GetCampaignByID(id primitive.ObjectID) (*model.CampaignResponse, error)
	GetAllCampaigns() ([]model.CampaignResponse, error)
	//UpdateCampaignAmount(id primitive.ObjectID, amount float64) error
	//DeleteCampaign(id primitive.ObjectID) error
}

type Campaign struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	CampaignID      string             `bson:"campaign_id"`
	Title           string             `bson:"title"`
	Description     string             `bson:"description"`
	Goal            float64            `bson:"goal"`
	CurrentAmount   float64            `bson:"current_amount"`
	CreatorID       primitive.ObjectID `bson:"creator"`
	CreatorUsername string             `bson:"creator_username"`
	LikeCount       int                `bson:"like_count"`
	CreatedAt       time.Time          `bson:"created_at"`
}
