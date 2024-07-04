package repository

import "context"

func (r *Repository) GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT name FROM test WHERE id = $1", input.Id).Scan(&output.Name)
	if err != nil {
		return
	}
	return
}

func (r *Repository) CreateEstate(id string, width, length int) error {
	_, err := r.Db.Exec("INSERT INTO estates (id, width, length) VALUES ($1, $2, $3)", id, width, length)
	return err
}

func (r *Repository) AddTree(id, estateId string, x, y, height int) error {
	_, err := r.Db.Exec("INSERT INTO trees (id, estate_id, x_coordinate, y_coordinate, height) VALUES ($1, $2, $3, $4, $5)", id, estateId, x, y, height)
	return err
}

func (r *Repository) GetEstateById(id string) (Estate, error) {
	var estate Estate
	err := r.Db.QueryRow("SELECT id, width, length FROM estates WHERE id = $1", id).Scan(&estate.Id, &estate.Width, &estate.Length)
	if err != nil {
		return estate, err
	}
	return estate, nil
}
