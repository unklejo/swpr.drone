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

	if request.Width <= 0 || request.Length <= 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Width and Length must be greater than 0"})
	}

	id := uuid.New().String()

	if err := s.Repository.CreateEstate(id, request.Width, request.Length); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create estate"})
	}

	return ctx.JSON(http.StatusCreated, map[string]string{"id": id})
}

func (s *Server) AddTreeToEstate(ctx echo.Context) error {
	estateId := ctx.Param("id")
	var request struct {
		X      int `json:"x"`
		Y      int `json:"y"`
		Height int `json:"height"`
	}

	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	treeId := uuid.New().String()
	if err := s.Repository.AddTree(treeId, estateId, request.X, request.Y, request.Height); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add tree"})
	}

	return ctx.JSON(http.StatusCreated, map[string]string{"id": treeId})
}
