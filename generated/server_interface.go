package generated

import (
	"github.com/labstack/echo/v4"
)

// ServerInterface is the interface that our server implementation will fulfill
type ServerInterface interface{}

// GetHelloParams represents the query parameters for the GetHello endpoint
type GetHelloParams struct {
	Name string `query:"name"`
	Id   int    `query:"id"`
}

// HelloResponse represents the response for the GetHello endpoint
type HelloResponse struct {
	Message string `json:"message"`
}

// RegisterHandlers registers all route handlers to the Echo instance
func RegisterHandlers(e *echo.Echo, server ServerInterface) {
	// e.GET("/hello", server.GetHelloWorld)
}
