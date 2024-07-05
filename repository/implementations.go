package repository

import (
	"context"
)

func (r *Repository) GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT name FROM test WHERE id = $1", input.Id).Scan(&output.Name)
	if err != nil {
		return
	}
	return
}

func (r *Repository) CreateEstate(id string, width, length int) (err error) {
	_, err = r.Db.Exec("INSERT INTO estates (id, width, length) VALUES ($1, $2, $3)", id, width, length)
	return err
}

func (r *Repository) AddTree(id, estateId string, x, y, height int) (err error) {
	_, err = r.Db.Exec("INSERT INTO trees (id, estate_id, x_coordinate, y_coordinate, height) VALUES ($1, $2, $3, $4, $5)", id, estateId, x, y, height)
	return err
}

func (r *Repository) GetEstateById(id string) (estate Estate, err error) {
	err = r.Db.QueryRow("SELECT id, width, length FROM estates WHERE id = $1", id).Scan(&estate.Id, &estate.Width, &estate.Length)
	if err != nil {
		return estate, err
	}
	return estate, nil
}

func (r *Repository) GetEstateStatsById(estateId string) (stats EstateStats, err error) {
	err = r.Db.QueryRow("SELECT COUNT(id), COALESCE(MAX(height), 0), COALESCE(MIN(height), 0), COALESCE(PERCENTILE_CONT(0.5) WITHIN GROUP (ORDER BY height), 0) FROM trees WHERE estate_id = $1", estateId).Scan(&stats.Count, &stats.MaxHeight, &stats.MinHeight, &stats.MedianHeight)
	if err != nil {
		return stats, err
	}

	return stats, nil
}

func (r *Repository) GetDronePlanByEstateId(estateId string) (distance int, err error) {
	err = r.Db.QueryRow("SELECT distance FROM drone_plans WHERE estate_id = $1", estateId).Scan(&distance)
	if err != nil {
		return 0, err
	}
	return distance, nil
}
