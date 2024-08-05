package infrastructureInterface

import (
	"context"
	"github.com/Many-Men/crowdfund_backend/internal/infrastructure/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user entity.User) (primitive.ObjectID, error)
	GetUserByID(ctx context.Context, id primitive.ObjectID) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	UpdateUserBalance(ctx context.Context, id primitive.ObjectID, balance float64) error
	DeleteUser(ctx context.Context, id primitive.ObjectID) error
	ListUsers(ctx context.Context) ([]entity.User, error)
}

type DonationRepository interface {
	CreateDonation(ctx context.Context, donation entity.Donation) (primitive.ObjectID, error)
	GetDonationByID(ctx context.Context, id primitive.ObjectID) (*entity.Donation, error)
	GetDonationsByCampaign(ctx context.Context, campaignID primitive.ObjectID) ([]entity.Donation, error)
	GetDonationsByDonor(ctx context.Context, donorID primitive.ObjectID) ([]entity.Donation, error)
	DeleteDonation(ctx context.Context, id primitive.ObjectID) error
}

type CampaignRepository interface {
	CreateCampaign(ctx context.Context, campaign entity.Campaign) (primitive.ObjectID, error)
	GetCampaignByID(ctx context.Context, id primitive.ObjectID) (*entity.Campaign, error)
	GetAllCampaigns(ctx context.Context) ([]entity.Campaign, error)
	UpdateCampaignAmount(ctx context.Context, id primitive.ObjectID, amount float64) error
	DeleteCampaign(ctx context.Context, id primitive.ObjectID) error
}
