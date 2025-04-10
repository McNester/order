package repositories

import (
	"errors"
	"orders/db"
	"orders/models"

	"github.com/jmoiron/sqlx"
)

type OrderRepo struct {
	db *sqlx.DB
}

func NewOrderRepo() *OrderRepo {
	return &OrderRepo{db: db.GetConnection()}
}

func (r *OrderRepo) ListOrders() ([]models.Order, error) {
	var orders []models.Order
	query := `
		SELECT *
		FROM orders
	`
	err := r.db.Select(&orders, query)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepo) GetOrder(id uint64) (*models.Order, error) {
	var order models.Order
	query := `
		SELECT *
		FROM orders
		WHERE id = ?
	`
	err := r.db.Get(&order, query, id)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepo) SaveOrder(order *models.Order) (*models.Order, error) {
	query := `
		INSERT INTO orders (status, quantity, product_id)
		VALUES (:status, :quantity, :product_id)
	`
	result, err := r.db.NamedExec(query, &order)
	if err != nil {
		return nil, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return r.GetOrder(uint64(lastId))
}

func (r *OrderRepo) UpdateOrder(id uint64, order *models.Order) (*models.Order, error) {
	query := `
    UPDATE orders SET 
    status = ?, quantity = ?, product_id = ?
    WHERE id = ?
    `
	res, err := r.db.Exec(
		query,
		order.Status,
		order.Quantity,
		order.ProductID,
		id)
	if err != nil {
		return nil, err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rows == 0 {
		return nil, errors.New("No order updated. Invalid ID?")
	}
	return r.GetOrder(id)
}

func (r *OrderRepo) DeleteOrder(id uint64) error {
	query := `
		DELETE FROM orders WHERE id = ?
	`
	res, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("No order deleted. Invalid ID?")
	}
	return nil
}
