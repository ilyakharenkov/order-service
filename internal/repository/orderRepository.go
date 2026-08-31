package repository

import (
	"database/sql"
	"errors"
	"log"
	"order-service/internal/repository/model"
)

type OrderRepository interface {
	FindAll() ([]model.Order, error)
	CreateOrder(order *model.Order) (*model.Order, error)
	FindOrder(id int) (model.Order, error)
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
	return repository.orders, nil
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

func (repository *orderRepositoryPostgres) FindOrder(id int) (model.Order, error) {
	for _, order := range repository.orders {
		if order.ID == id {
			return order, nil
		}
	}
	return model.Order{}, errors.New("order not found")
}
