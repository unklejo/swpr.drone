package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// Handler for the /estate endpoint
func (s *Server) CreateEstate(ctx echo.Context) error {
	var request struct {
		Width  int `json:"width"`
		Length int `json:"length"`
	}
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	id := uuid.New().String()

	if err := s.Repository.CreateEstate(id, request.Width, request.Length); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create estate"})
	}

	return ctx.JSON(http.StatusCreated, map[string]string{"id": id})
}
