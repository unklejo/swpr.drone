package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/lib/pq"
	"github.com/unklejo/swpr.drone/repository"
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

	// Error handling regarding database and foreign key
	treeId := uuid.New().String()
	err := s.Repository.AddTree(treeId, estateId, request.X, request.Y, request.Height)
	if err != nil {
		if err == repository.ErrForeignKeyNotFound {
			return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Related resource not found"})
		}

		if err, ok := err.(*pq.Error); ok && err.Code == "23505" { // Unique violation
			return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Plot already has a tree"})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add tree"})
	}

	return ctx.JSON(http.StatusCreated, map[string]string{"id": treeId})
}
