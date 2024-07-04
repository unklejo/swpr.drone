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
