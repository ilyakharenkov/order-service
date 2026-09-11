package repository

import (
	"database/sql"
	"fmt"
	"log"
	"order-service/internal/repository/model"
	"time"
)

type OrderRepository interface {
	FindAll() ([]model.Order, error)
	CreateOrder(order *model.Order) (*model.Order, error)
	FindOrder(orderNumber string) (model.Order, error)
	CancelOrder(orderNumber string) (model.Order, error)
}

type orderRepositoryPostgres struct {
	db     *sql.DB
	orders []model.Order
}

func NewOrderRepository(db *sql.DB, orders []model.Order) OrderRepository {
	return &orderRepositoryPostgres{
		db:     db,
		orders: orders,
	}
}

func (repository *orderRepositoryPostgres) FindAll() ([]model.Order, error) {
	query := "SELECT id, order_number, sku, quantity, status, created_at, updated_at FROM order_t"
	rows, err := repository.db.Query(query)
	if err != nil {
		return nil, err
	}

	var orders []model.Order
	for rows.Next() {
		var order model.Order
		err := rows.Scan(
			&order.ID,
			&order.OrderNumber,
			&order.SKU,
			&order.Quantity,
			&order.Status,
			&order.CreatedAt,
			&order.UpdatedAt)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, rows.Err()
}

func (repository *orderRepositoryPostgres) CreateOrder(order *model.Order) (*model.Order, error) {
	query := "INSERT INTO order_t (order_number, sku, quantity, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"

	err := repository.db.QueryRow(query,
		order.OrderNumber,
		order.SKU,
		order.Quantity,
		order.Status,
		order.CreatedAt,
		order.UpdatedAt,
	).Scan(&order.ID)

	if err != nil {
		log.Printf("Failed to create product: %v", err)
		return nil, err
	}

	return order, nil
}

func (repository *orderRepositoryPostgres) FindOrder(orderNumber string) (model.Order, error) {
	qeury := "SELECT id, order_number, sku, quantity, status, created_at, updated_at FROM order_t WHERE order_t.order_number = $1"
	row := repository.db.QueryRow(qeury, orderNumber)
	var order model.Order
	err := row.Scan(
		&order.ID,
		&order.OrderNumber,
		&order.SKU,
		&order.Quantity,
		&order.Status,
		&order.CreatedAt,
		&order.UpdatedAt)
	if err != nil {
		return model.Order{}, err
	}

	return order, nil
}

func (repository *orderRepositoryPostgres) CancelOrder(orderNumber string) (model.Order, error) {
	query := "UPDATE order_t SET status = 'CANCELLED' WHERE order_t.id IN (SELECT order_t.id FROM order_t WHERE order_t.id = $1) RETURNING order_t.id"

	row := repository.db.QueryRow(query, orderNumber)

	fmt.Println(row)

	return model.Order{
		OrderNumber: orderNumber,
		Status:      "CANCELLED",
		SKU:         "",
		Quantity:    0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
