package generated

import (
	"github.com/labstack/echo/v4"
)

// Interface that our server implementation will fulfill
type ServerInterface interface {
	CreateEstate(ctx echo.Context) error
	AddTreeToEstate(ctx echo.Context) error
	GetEstateStats(ctx echo.Context) error
	GetDronePlan(ctx echo.Context) error
}

// Registers all route handlers to the Echo instance
func RegisterHandlers(e *echo.Echo, server ServerInterface) {
	e.POST("/estate", server.CreateEstate)
	e.POST("/estate/:id/tree", server.AddTreeToEstate)
	e.GET("/estate/:id/stats", server.GetEstateStats)
	e.GET("/estate/:id/drone-plan", server.GetDronePlan)
}
