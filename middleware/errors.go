package _middleware

import (
	"github.com/Many-Men/crowdfund_backend/utils"
	"github.com/labstack/echo/v4"
)

func ErrorHandlingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			statusCode := utils.MapErrorToStatusCode(err)
			return c.JSON(statusCode, map[string]string{"error": err.Error()})
		}
		return nil
	}
}
