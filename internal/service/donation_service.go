package service

import (
	"context"
	"github.com/Many-Men/crowdfund_backend/config"
	_interface "github.com/Many-Men/crowdfund_backend/internal/domain/interface"
	"github.com/Many-Men/crowdfund_backend/internal/infrastructure/entity"
	infrastructureInterface "github.com/Many-Men/crowdfund_backend/internal/service/interface"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DonationServiceImpl struct {
	donationRepository infrastructureInterface.DonationRepository
	config             *config.Config
}

func NewDonationServiceImpl(donationRepo infrastructureInterface.DonationRepository) _interface.DonationService {
	return &DonationServiceImpl{
		donationRepository: donationRepo,
		config:             config.Load(),
	}
}

func (s *DonationServiceImpl) CreateDonation(donation entity.Donation) (primitive.ObjectID, error) {
	return s.donationRepository.CreateDonation(context.Background(), donation)
}

func (s *DonationServiceImpl) GetDonationByID(id primitive.ObjectID) (*entity.Donation, error) {
	return s.donationRepository.GetDonationByID(context.Background(), id)
}

func (s *DonationServiceImpl) GetDonationsByCampaign(campaignID primitive.ObjectID) ([]entity.Donation, error) {
	return s.donationRepository.GetDonationsByCampaign(context.Background(), campaignID)
}

func (s *DonationServiceImpl) GetDonationsByDonor(donorID primitive.ObjectID) ([]entity.Donation, error) {
	return s.donationRepository.GetDonationsByDonor(context.Background(), donorID)
}

func (s *DonationServiceImpl) DeleteDonation(id primitive.ObjectID) error {
	return s.donationRepository.DeleteDonation(context.Background(), id)
}
