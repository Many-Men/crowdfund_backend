package delivery

import (
	"github.com/Many-Men/crowdfund_backend/internal/delivery/controller"
	"github.com/Many-Men/crowdfund_backend/internal/infrastructure/repository"
	"github.com/Many-Men/crowdfund_backend/internal/service"
	_middleware "github.com/Many-Men/crowdfund_backend/middleware"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes(e *echo.Echo, db *mongo.Database) {
	ur := repository.NewUserRepositoryImpl(db)
	cr := repository.NewCampaignRepositoryImpl(db)
	dr := repository.NewDonationRepositoryImpl(db)

	us := service.NewUserServiceImpl(ur)
	cs := service.NewCampaignServiceImpl(cr)
	ds := service.NewDonationServiceImpl(dr)

	c := controller.NewAppController(us, ds, cs)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/user", c.CreateUser)
	e.GET("/user/:id", c.GetUserByID)
	e.PUT("/user/:id/balance", c.UpdateUserBalance)
	e.DELETE("/user/:id", c.DeleteUser)
	e.GET("/users", c.ListUsers)

	e.POST("/donation", c.CreateDonation)
	e.GET("/donation/:id", c.GetDonationByID)
	e.GET("/donations/campaign/:campaign_id", c.GetDonationsByCampaign)
	e.GET("/donations/donor/:donor_id", c.GetDonationsByDonor)
	e.DELETE("/donation/:id", c.DeleteDonation)

	e.POST("/campaign", c.CreateCampaign, _middleware.ValidateAccessTokenMiddleware())
	e.GET("/campaign/:id", c.GetCampaignByID)
	e.GET("/campaigns", c.GetAllCampaigns, _middleware.ValidateAccessTokenMiddleware())
	e.PUT("/campaign/:id/amount", c.UpdateCampaignAmount)
	e.DELETE("/campaign/:id", c.DeleteCampaign)
}
