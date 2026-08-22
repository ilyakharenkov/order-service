package repository

import (
	"errors"
	"order-service/internal/repository/model"
)

type OrderRepository interface {
	FindAll() ([]model.Order, error)
	CreateOrder(order model.Order) (model.Order, error)
	FindOrder(id int) (model.Order, error)
}

type orderRepositoryPostgres struct {
	orders []model.Order
}

func NewOrderRepository(orders []model.Order) OrderRepository {
	return &orderRepositoryPostgres{
		orders: orders,
	}
}

func (repository *orderRepositoryPostgres) FindAll() ([]model.Order, error) {
	return repository.orders, nil
}

func (repository *orderRepositoryPostgres) CreateOrder(order model.Order) (model.Order, error) {
	repository.orders = append(repository.orders, order)
	return order, errors.New("errors")
}

func (repository *orderRepositoryPostgres) FindOrder(id int) (model.Order, error) {
	for _, order := range repository.orders {
		if order.ID == id {
			return order, nil
		}
	}
	return model.Order{}, errors.New("order not found")
}
