package _interface

import (
	"github.com/Many-Men/crowdfund_backend/internal/delivery/model"
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
	CreateCampaign(title, description, username string, goal float64, pictures [][]byte) error
	//GetCampaignByID(id primitive.ObjectID) (*model.CampaignResponse, error)
	GetAllCampaigns() ([]model.CampaignResponse, error)
	//UpdateCampaignAmount(id primitive.ObjectID, amount float64) error
	//DeleteCampaign(id primitive.ObjectID) error
}

type FileService interface {
	SaveFile(filename string, content []byte) error
	LoadFile(filename string) ([]byte, error)
	UpdateFileName(oldName, newName string) error
	UpdateFile(newFileBytes []byte, fileName string) error
	DeleteFile(filename string) error
}
