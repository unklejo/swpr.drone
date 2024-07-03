package generated

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ServerInterface is the interface that our server implementation will fulfill
type ServerInterface interface {
	GetHelloWorld(ctx echo.Context) error
}

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
	e.GET("/hello", func(ctx echo.Context) error {
		params := GetHelloParams{}
		if err := ctx.Bind(&params); err != nil {
			return err
		}
		response := HelloResponse{Message: "Hello, " + params.Name + "!"}
		return ctx.JSON(http.StatusOK, response)
	})
}
