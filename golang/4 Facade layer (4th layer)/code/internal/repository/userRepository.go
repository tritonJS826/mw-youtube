package repository

import (
	"database/sql"
	"facade/model"
)

type UserRepository struct {
	db *sql.DB
}

func (r *UserRepository) FindByID(id string) (*model.User, error) {

	var user model.User
	err := r.db.QueryRow("SELECT id, name FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
