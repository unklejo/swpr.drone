package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/lib/pq"
)

// 1. Handler for POST `/estate` endpoint
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

// 2. Handler for POST `/estate/:id/tree` endpoint
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

	// Tree height is 0 or lower
	if request.Height < 0 || request.Height >= 30 {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Height must be within 1 to 30 meters"})
	}

	// Check the estate exist or not
	estate, err := s.Repository.GetEstateById(estateId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Related resource not found"})
		}
	}

	// Coordinates out of bounds from estate's plot
	if request.X < 0 || request.Y < 0 || request.X >= estate.Width || request.Y >= estate.Length {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Coordinates out of bounds"})
	}

	// Error handling regarding database and foreign key
	treeId := uuid.New().String()
	err = s.Repository.AddTree(treeId, estateId, request.X, request.Y, request.Height)
	if err != nil {
		// Tree already exists in the plot (handling racing condition)
		if err, ok := err.(*pq.Error); ok && err.Code == "23505" { // Unique violation
			return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Plot already has a tree"})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add tree"})
	}

	return ctx.JSON(http.StatusCreated, map[string]string{"id": treeId})
}

// 3. Handler for GET `/estate/:id/stats` endpoint
func (s *Server) GetEstateStats(ctx echo.Context) error {
	estateId := ctx.Param("id")
	stats, err := s.Repository.GetEstateStatsById(estateId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Estate not found"})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve estate stats"})
	}

	return ctx.JSON(http.StatusOK, stats)
}
