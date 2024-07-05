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
func (s *Server) PostEstate(ctx echo.Context) error {
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

	id, err := s.Repository.CreateEstate(request.Width, request.Length)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create estate"})
	}

	return ctx.JSON(http.StatusCreated, map[string]string{"id": id})
}

// 2. Handler for POST `/estate/:id/tree` endpoint
func (s *Server) PostEstateIdTree(ctx echo.Context, uuid uuid.UUID) error {
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
			return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Estate not found"})
		}
	}

	// Coordinates out of bounds from estate's plot
	if request.X <= 0 || request.Y <= 0 || request.X > estate.Width || request.Y > estate.Length {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Coordinates out of bounds"})
	}

	// Error handling regarding database and foreign key
	id, err := s.Repository.AddTree(estateId, request.X, request.Y, request.Height)
	if err != nil {
		// Tree already exists in the plot (handling racing condition)
		if err, ok := err.(*pq.Error); ok && err.Code == "23505" { // Unique violation
			return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Plot already has a tree"})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add tree"})
	}

	return ctx.JSON(http.StatusCreated, map[string]string{"id": id})
}

// 3. Handler for GET `/estate/:id/stats` endpoint
func (s *Server) GetEstateIdStats(ctx echo.Context, uuid uuid.UUID) error {
	estateId := ctx.Param("id")

	// Check the estate exist or not, just like in AddTree
	_, err := s.Repository.GetEstateById(estateId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Estate not found"})
		}
	}

	stats, err := s.Repository.GetEstateStatsById(estateId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve estate stats"})
	}

	return ctx.JSON(http.StatusOK, stats)
}

// 4. Handler for GET `/estate/:id/drone-plan` endpoint
func (s *Server) GetEstateIdDronePlan(ctx echo.Context, uuid uuid.UUID) error {
	estateId := ctx.Param("id")

	// Check the estate exist or not, just like in AddTree
	_, err := s.Repository.GetEstateById(estateId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Estate not found"})
		}
	}

	distance, err := s.Repository.GetDronePlanByEstateId(estateId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, map[string]string{"error": "Drone plan not found"})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve drone plans"})
	}

	return ctx.JSON(http.StatusOK, distance)
}

// Get coordinate
func getCoordinate(maxDistance int) error {
	i := 0

	for i < maxDistance {
		// 1. check next post
		// 1.1. if post not empty
		// 1.1.1. traverse to next x_coordinate
		// 1.1.2. else check north
		// 1.1.1.1. if exist: traverse to next y_coordinate
		// 1.1.1.2. else: end
	}
	return nil
}
