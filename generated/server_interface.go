package generated

import (
	"github.com/labstack/echo/v4"
)

// ServerInterface is the interface that our server implementation will fulfill
type ServerInterface interface {
	CreateEstate(ctx echo.Context) error
}

// RegisterHandlers registers all route handlers to the Echo instance
func RegisterHandlers(e *echo.Echo, server ServerInterface) {
	e.POST("/estate", server.CreateEstate)
}
