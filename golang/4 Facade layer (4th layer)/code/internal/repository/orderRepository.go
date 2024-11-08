package repository

import (
	"database/sql"
	"facade/model"
)

type OrderRepository struct {
	db *sql.DB
}

func (r *OrderRepository) FindByID(id string) (*model.Order, error) {

	var order model.Order
	err := r.db.QueryRow("SELECT id, order FROM orders WHERE id = ?", id).Scan(&order.ID, &order.Order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
