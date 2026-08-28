package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"order-service/configs"
	"order-service/internal/handlers"
	"order-service/internal/repository"
	"order-service/internal/repository/model"
	"order-service/internal/service"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	postgresConfig := configs.PostgresConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		postgresConfig.DBHost, postgresConfig.DBPort, postgresConfig.DBUser, postgresConfig.DBPassword, postgresConfig.DBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	orders := make([]model.Order, 0, 10)
	now := time.Now()

	statuses := []model.OrderStatus{
		model.StatusPending,
		model.StatusConfirmed,
		model.StatusCancelled,
		model.StatusFailed,
		model.StatusPending,
		model.StatusConfirmed,
		model.StatusPending,
		model.StatusConfirmed,
		model.StatusCancelled,
		model.StatusFailed,
	}

	for i := 0; i < 10; i++ {
		order := model.Order{
			ID:          i + 1,
			OrderNumber: fmt.Sprintf("ORD-%03d", i+1),
			SKU:         fmt.Sprintf("SKU-%04d", 1001+i),
			Quantity:    (i+1)*2 + 3,
			Status:      statuses[i],
			CreatedAt:   now.Add(-time.Duration(10-i) * time.Hour),
			UpdatedAt:   now.Add(-time.Duration(10-i-1) * time.Hour),
		}

		orders = append(orders, order)
	}

	orderRepository := repository.NewOrderRepository(db, orders)
	orderService := service.NewOrderService(orderRepository)
	orderHandler := handlers.NewOrderHttpHandler(orderService)

	http.HandleFunc("GET /orders", orderHandler.FindAll)

	if err := http.ListenAndServe("localhost:8081", nil); err != nil {
		fmt.Println(err)
	}
}
