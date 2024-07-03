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

// type Server struct{}

// func newServer() *Server {
// 	return &Server{}
// }

func (s *Server) GetHelloWorld(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}
