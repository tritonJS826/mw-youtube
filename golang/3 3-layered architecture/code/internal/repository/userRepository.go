package repository

import (
	"database/sql"
	"import-swag/internal/service"
)

type UserRepository struct {
	db *sql.DB
}

func (r *UserRepository) FindByID(id string) (*service.User, error) {

	var user User
	err := r.db.QueryRow("SELECT id, name FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
