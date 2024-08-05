package server

import (
	"github.com/Many-Men/crowdfund_backend/config"
	"github.com/Many-Men/crowdfund_backend/internal/delivery"
	_middleware "github.com/Many-Men/crowdfund_backend/middleware"
	"go.mongodb.org/mongo-driver/mongo"

	_ "github.com/Many-Men/crowdfund_backend/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"os"
)

// RunHTTPServer
// @title ...
// @version 1.0
// @description ...
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name .
// @license.url .
// @host petstore.swagger.io
// @externalDocs.description  OpenAPI 2.0
// @BasePath /
func RunHTTPServer(cfg *config.Config, db *mongo.Database) {
	e, log := echo.New(), logrus.New()

	log.Out = os.Stdout
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} [${status}] ${method} ${uri} (${latency_human})\n",
		Output: log.Out,
	}))
	e.Use(middleware.CORS())
	e.Use(_middleware.ErrorHandlingMiddleware)

	delivery.RegisterRoutes(e, db)
	if err := e.Start(cfg.Server.Port); err != nil {
		panic(err)
	}
}
