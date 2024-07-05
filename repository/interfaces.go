// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import "context"

type Estate struct {
	Id     string
	Width  int
	Length int
}

type EstateStats struct {
	Count        int `json:"count"`
	MaxHeight    int `json:"max"`
	MinHeight    int `json:"min"`
	MedianHeight int `json:"median"`
}

type DronePlan struct {
	Distance int `json:"distance"`
}

type RepositoryInterface interface {
	GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error)
	CreateEstate(width, length int) (id string, err error)
	AddTree(estateId string, x, y, height int) (id string, err error)
	GetEstateById(id string) (estate Estate, err error)
	GetEstateStatsById(estateId string) (stats EstateStats, err error)
	GetDronePlanByEstateId(estateId string) (plan DronePlan, err error)
}
