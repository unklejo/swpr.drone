package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/unklejo/swpr.drone/repository"
)

type Server struct {
	Repository repository.RepositoryInterface
}

type NewServerOptions struct {
	Repository repository.RepositoryInterface
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{}
}

func (s *Server) GetHelloWorld(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}

// func (s *Server) GetHelloWorld(ctx echo.Context) error {
// 	params := new(generated.GetHelloParams)
// 	if err := ctx.Bind(params); err != nil {
// 		return err
// 	}
// 	response := generated.HelloResponse{Message: "Hello, " + params.Name + "!"}
// 	return ctx.JSON(http.StatusOK, response)
// }
