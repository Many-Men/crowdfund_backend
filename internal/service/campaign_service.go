package service

import (
	"context"
	"github.com/Many-Men/crowdfund_backend/config"
	"github.com/Many-Men/crowdfund_backend/internal/delivery/model"
	_interface "github.com/Many-Men/crowdfund_backend/internal/domain/interface"
	"github.com/Many-Men/crowdfund_backend/internal/infrastructure/entity"
	infrastructureInterface "github.com/Many-Men/crowdfund_backend/internal/service/interface"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CampaignServiceImpl struct {
	campaignRepository infrastructureInterface.CampaignRepository
	fileService        _interface.FileService
	config             *config.Config
}

func NewCampaignServiceImpl(campaignRepo infrastructureInterface.CampaignRepository, fileService _interface.FileService) _interface.CampaignService {
	return &CampaignServiceImpl{
		campaignRepository: campaignRepo,
		fileService:        fileService,
		config:             config.Load(),
	}
}

func (s *CampaignServiceImpl) CreateCampaign(title, description, username string, goal float64, pictures [][]byte) error {
	user, err := s.campaignRepository.GetUserByUsername(context.Background(), username)
	if err != nil {
		return err
	}

	id := uuid.New().String()
	content := make([]string, len(pictures))

	for i, p := range pictures {
		fn := id + "_" + string(rune(i))
		content = append(content, fn)

		go s.fileService.SaveFile(fn, p)
	}

	campaign := entity.NewCampaign(title, description, username, id, goal, user.ID, content)
	if _, err := s.campaignRepository.CreateCampaign(context.Background(), *campaign); err != nil {
		return err
	}

	return nil
}

func (s *CampaignServiceImpl) GetCampaignByID(id primitive.ObjectID) (*entity.Campaign, error) {
	return s.campaignRepository.GetCampaignByID(context.Background(), id)
}

func (s *CampaignServiceImpl) GetAllCampaigns() ([]model.CampaignResponse, error) {
	campaigns, err := s.campaignRepository.GetAllCampaigns(context.Background())
	if err != nil {
		return nil, err
	}

	var campaignResponses []model.CampaignResponse
	for _, campaign := range campaigns {
		campaignResponses = append(campaignResponses, model.CampaignResponse{
			CampaignID:      campaign.CampaignID,
			Title:           campaign.Title,
			Description:     campaign.Description,
			Goal:            campaign.Goal,
			CurrentAmount:   campaign.CurrentAmount,
			CreatorUsername: campaign.CreatorUsername,
			LikeCount:       campaign.LikeCount,
			CreatedAt:       campaign.CreatedAt,
		})
	}

	return campaignResponses, nil
}

func (s *CampaignServiceImpl) UpdateCampaignAmount(id primitive.ObjectID, amount float64) error {
	return s.campaignRepository.UpdateCampaignAmount(context.Background(), id, amount)
}

func (s *CampaignServiceImpl) DeleteCampaign(id primitive.ObjectID) error {
	return s.campaignRepository.DeleteCampaign(context.Background(), id)
}
