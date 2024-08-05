package service

import (
	"context"
	"github.com/Many-Men/crowdfund_backend/config"
	_interface "github.com/Many-Men/crowdfund_backend/internal/domain/interface"
	"github.com/Many-Men/crowdfund_backend/internal/infrastructure/entity"
	infrastructureInterface "github.com/Many-Men/crowdfund_backend/internal/service/interface"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CampaignServiceImpl struct {
	campaignRepository infrastructureInterface.CampaignRepository
	config             *config.Config
}

func NewCampaignServiceImpl(campaignRepo infrastructureInterface.CampaignRepository) _interface.CampaignService {
	return &CampaignServiceImpl{
		campaignRepository: campaignRepo,
		config:             config.Load(),
	}
}

func (s *CampaignServiceImpl) CreateCampaign(campaign entity.Campaign) (primitive.ObjectID, error) {
	return s.campaignRepository.CreateCampaign(context.Background(), campaign)
}

func (s *CampaignServiceImpl) GetCampaignByID(id primitive.ObjectID) (*entity.Campaign, error) {
	return s.campaignRepository.GetCampaignByID(context.Background(), id)
}

func (s *CampaignServiceImpl) GetAllCampaigns() ([]entity.Campaign, error) {
	return s.campaignRepository.GetAllCampaigns(context.Background())
}

func (s *CampaignServiceImpl) UpdateCampaignAmount(id primitive.ObjectID, amount float64) error {
	return s.campaignRepository.UpdateCampaignAmount(context.Background(), id, amount)
}

func (s *CampaignServiceImpl) DeleteCampaign(id primitive.ObjectID) error {
	return s.campaignRepository.DeleteCampaign(context.Background(), id)
}
