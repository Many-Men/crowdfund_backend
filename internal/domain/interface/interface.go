package _interface

import (
	"github.com/Many-Men/crowdfund_backend/internal/infrastructure/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	CreateUser(user entity.User) (primitive.ObjectID, error)
	GetUserByID(id primitive.ObjectID) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	UpdateUserBalance(id primitive.ObjectID, balance float64) error
	DeleteUser(id primitive.ObjectID) error
	ListUsers() ([]entity.User, error)
}

type DonationService interface {
	CreateDonation(donation entity.Donation) (primitive.ObjectID, error)
	GetDonationByID(id primitive.ObjectID) (*entity.Donation, error)
	GetDonationsByCampaign(campaignID primitive.ObjectID) ([]entity.Donation, error)
	GetDonationsByDonor(donorID primitive.ObjectID) ([]entity.Donation, error)
	DeleteDonation(id primitive.ObjectID) error
}

type CampaignService interface {
	CreateCampaign(campaign entity.Campaign) (primitive.ObjectID, error)
	GetCampaignByID(id primitive.ObjectID) (*entity.Campaign, error)
	GetAllCampaigns() ([]entity.Campaign, error)
	UpdateCampaignAmount(id primitive.ObjectID, amount float64) error
	DeleteCampaign(id primitive.ObjectID) error
}
