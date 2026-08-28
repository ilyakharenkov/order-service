package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"order-service/internal/service"
	"order-service/internal/service/dto"
)

type OrderHandler interface {
	FindAll(w http.ResponseWriter, r *http.Request)
	CreateOrder(w http.ResponseWriter, r *http.Request)
}

type orderHttpHandler struct {
	service service.OrderService
}

func NewOrderHttpHandler(service service.OrderService) OrderHandler {
	return &orderHttpHandler{
		service: service,
	}
}

func (handler *orderHttpHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	response, err1 := handler.service.FindAll()
	if err1 != nil {
		log.Printf("Error %v", err1)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err2 := json.NewEncoder(w).Encode(response)
	if err2 != nil {
		log.Printf("Error %v", err2)
		return
	}
}

func (handler *orderHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	requestBody := &dto.Order{}
	if err := json.NewDecoder(r.Body).Decode(requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := handler.service.CreateOrder(requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err2 := json.NewEncoder(w).Encode(response)
	if err2 != nil {
		log.Printf("Error %v", err2)
		return
	}
}
