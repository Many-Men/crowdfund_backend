package delivery

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RegisterRoutes(e *echo.Echo, db *sqlx.DB) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
