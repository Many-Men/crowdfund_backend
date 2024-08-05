package controller

import (
	"github.com/Many-Men/crowdfund_backend/config"
	_errors "github.com/Many-Men/crowdfund_backend/errors"
	"github.com/Many-Men/crowdfund_backend/internal/delivery/model"
	_interface "github.com/Many-Men/crowdfund_backend/internal/domain/interface"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AppController struct {
	userService     _interface.UserService
	donationService _interface.DonationService
	campaignService _interface.CampaignService
	config          *config.Config
}

func NewAppController(userService _interface.UserService, donationService _interface.DonationService, campaignService _interface.CampaignService) *AppController {
	return &AppController{
		userService:     userService,
		donationService: donationService,
		campaignService: campaignService,
		config:          config.Load(),
	}
}

func (ac *AppController) CreateUser(c echo.Context) error {
	var request model.UserRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
	}

	if err := ac.userService.CreateUser(request.Username); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, model.SuccessResponse{Message: "User created successfully"})
}

//
//func (ac *AppController) GetUserByID(c echo.Context) error {
//	id, err := primitive.ObjectIDFromHex(c.Param("id"))
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid user ID"})
//	}
//
//	user, err := ac.userService.GetUserByID(id)
//	if err != nil {
//		return err
//	}
//
//	return c.JSON(http.StatusOK, user)
//}
//
//func (ac *AppController) UpdateUserBalance(c echo.Context) error {
//	id, err := primitive.ObjectIDFromHex(c.Param("id"))
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid user ID"})
//	}
//
//	var request model.UserUpdateBalanceRequest
//	if err := c.Bind(&request); err != nil {
//		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
//	}
//
//	if err = ac.userService.UpdateUserBalance(id, request.Balance); err != nil {
//		return err
//	}
//
//	return c.JSON(http.StatusOK, model.SuccessResponse{Message: "User balance updated successfully"})
//}
//
//func (ac *AppController) DeleteUser(c echo.Context) error {
//	id, err := primitive.ObjectIDFromHex(c.Param("id"))
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid user ID"})
//	}
//
//	if err = ac.userService.DeleteUser(id); err != nil {
//		return err
//	}
//
//	return c.JSON(http.StatusOK, model.SuccessResponse{Message: "User deleted successfully"})
//}
//
//func (ac *AppController) ListUsers(c echo.Context) error {
//	users, err := ac.userService.ListUsers()
//	if err != nil {
//		return err
//	}
//
//	return c.JSON(http.StatusOK, users)
//}

//func (ac *AppController) CreateDonation(c echo.Context) error {
//	var request model.DonationRequest
//	if err := c.Bind(&request); err != nil {
//		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
//	}
//
//	donation := entity.Donation{
//		Campaign: request.CampaignID,
//		Donor:    request.DonorID,
//		Amount:   request.Amount,
//	}
//
//	if _, err := ac.donationService.CreateDonation(donation); err != nil {
//		return err
//	}
//
//	return c.JSON(http.StatusCreated, model.SuccessResponse{Message: "Donation created successfully"})
//}
//
//func (ac *AppController) GetDonationByID(c echo.Context) error {
//	id, err := primitive.ObjectIDFromHex(c.Param("id"))
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid donation ID"})
//	}
//
//	donation, err := ac.donationService.GetDonationByID(id)
//	if err != nil {
//		return err
//	}
//
//	return c.JSON(http.StatusOK, donation)
//}
//
//func (ac *AppController) GetDonationsByCampaign(c echo.Context) error {
//	campaignID, err := primitive.ObjectIDFromHex(c.Param("campaign_id"))
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid campaign ID"})
//	}
//
//	donations, err := ac.donationService.GetDonationsByCampaign(campaignID)
//	if err != nil {
//		return err
//	}
//
//	return c.JSON(http.StatusOK, donations)
//}

//func (ac *AppController) GetDonationsByDonor(c echo.Context) error {
//	donorID, err := primitive.ObjectIDFromHex(c.Param("donor_id"))
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid donor ID"})
//	}
//
//	donations, err := ac.donationService.GetDonationsByDonor(donorID)
//	if err != nil {
//		return err
//	}
//
//	return c.JSON(http.StatusOK, donations)
//}
//
//func (ac *AppController) DeleteDonation(c echo.Context) error {
//	id, err := primitive.ObjectIDFromHex(c.Param("id"))
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid donation ID"})
//	}
//
//	if err = ac.donationService.DeleteDonation(id); err != nil {
//		return err
//	}
//
//	return c.JSON(http.StatusOK, model.SuccessResponse{Message: "Donation deleted successfully"})
//}

func (ac *AppController) CreateCampaign(c echo.Context) error {
	var request model.CampaignRequest
	if err := c.Bind(&request); err != nil {
		return &_errors.BadRequestError{Message: "invalid request"}
	}

	username := c.Request().Context().Value("username").(string)
	if err := ac.campaignService.CreateCampaign(request.Title, request.Description, username, request.GoalAmount, request.Pictures); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, model.SuccessResponse{Message: "Campaign created successfully"})
}

//func (ac *AppController) GetCampaignByID(c echo.Context) error {
//	id, err := primitive.ObjectIDFromHex(c.Param("id"))
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid campaign ID"})
//	}
//
//	campaign, err := ac.campaignService.GetCampaignByID(id)
//	if err != nil {
//		return err
//	}
//
//	return c.JSON(http.StatusOK, campaign)
//}

func (ac *AppController) GetAllCampaigns(c echo.Context) error {
	campaigns, err := ac.campaignService.GetAllCampaigns()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, campaigns)
}

//func (ac *AppController) UpdateCampaignAmount(c echo.Context) error {
//	id, err := primitive.ObjectIDFromHex(c.Param("id"))
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid campaign ID"})
//	}
//
//	var request struct {
//		Amount float64 `json:"amount"`
//	}
//	if err := c.Bind(&request); err != nil {
//		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
//	}
//
//	if err = ac.campaignService.UpdateCampaignAmount(id, request.Amount); err != nil {
//		return err
//	}
//
//	return c.JSON(http.StatusOK, model.SuccessResponse{Message: "Campaign amount updated successfully"})
//}
//
//func (ac *AppController) DeleteCampaign(c echo.Context) error {
//	id, err := primitive.ObjectIDFromHex(c.Param("id"))
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid campaign ID"})
//	}
//
//	if err = ac.campaignService.DeleteCampaign(id); err != nil {
//		return err
//	}
//
//	return c.JSON(http.StatusOK, model.SuccessResponse{Message: "Campaign deleted successfully"})
//}
