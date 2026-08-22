package service

import (
	"fmt"
	"order-service/internal/repository"
	"order-service/internal/repository/model"
	"order-service/internal/service/dto"
)

type OrderService interface {
	FindAll() ([]dto.Order, error)
	CreateOrder(order dto.Order) (dto.Order, error)
	FindOrder(id int) (dto.Order, error)
}

type orderService struct {
	repository repository.OrderRepository
}

func NewOrderService(repository repository.OrderRepository) OrderService {
	return &orderService{
		repository: repository,
	}
}

func (service *orderService) FindAll() ([]dto.Order, error) {
	orders, err := service.repository.FindAll()
	if err != nil {
		return nil, err
	}

	fmt.Println(orders)

	return nil, err
}

func (service *orderService) CreateOrder(order dto.Order) (dto.Order, error) {
	createOrder, err := service.repository.CreateOrder(model.Order{})
	if err != nil {
		return dto.Order{}, err
	}
	fmt.Println(createOrder)

	return dto.Order{}, err
}

func (service *orderService) FindOrder(id int) (dto.Order, error) {
	order, err := service.repository.FindOrder(id)
	if err != nil {
		return dto.Order{}, err
	}

	fmt.Println(order)

	return dto.Order{}, err
}
