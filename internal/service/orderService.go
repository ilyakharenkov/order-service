package service

import (
	"fmt"
	"order-service/internal/repository"
	"order-service/internal/repository/model"
	"order-service/internal/service/dto"
	"time"
)

type OrderService interface {
	FindAll() ([]dto.Order, error)
	CreateOrder(order *dto.Order) (*dto.Order, error)
	FindOrder(id string) (*dto.Order, error)
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

	var sliceOrder = make([]dto.Order, 0, len(orders))
	for _, it := range orders {
		o1 := dto.Order{
			OrderNumber: it.OrderNumber,
			SKU:         it.SKU,
			Quantity:    it.Quantity,
			Status:      dto.OrderStatus(it.Status),
			CreatedAt:   it.CreatedAt,
			UpdatedAt:   it.UpdatedAt,
		}

		sliceOrder = append(sliceOrder, o1)
	}

	return sliceOrder, err
}

func (service *orderService) CreateOrder(order *dto.Order) (*dto.Order, error) {
	createOrder, err := service.repository.CreateOrder(&model.Order{
		OrderNumber: order.OrderNumber,
		SKU:         order.SKU,
		Quantity:    order.Quantity,
		Status:      model.OrderStatus(order.Status),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	})

	if err != nil {
		return nil, err
	}
	fmt.Println(createOrder)

	return &dto.Order{
		OrderNumber: createOrder.OrderNumber,
		SKU:         createOrder.SKU,
		Quantity:    createOrder.Quantity,
		Status:      dto.OrderStatus(createOrder.Status),
		CreatedAt:   createOrder.CreatedAt,
		UpdatedAt:   createOrder.UpdatedAt,
	}, err
}

func (service *orderService) FindOrder(id string) (*dto.Order, error) {
	order, err := service.repository.FindOrder(id)
	if err != nil {
		return &dto.Order{}, err
	}

	return &dto.Order{
		OrderNumber: order.OrderNumber,
		SKU:         order.SKU,
		Quantity:    order.Quantity,
		Status:      dto.OrderStatus(order.Status),
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
	}, err
}
