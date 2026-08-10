order-service/
├── cmd/
│   └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── domain/
│   │   └── order.go
│   ├── repository/
│   │   ├── postgres.go
│   │   └── order_repo.go
│   ├── service/
│   │   └── order_service.go
│   ├── handlers/
│   │   └── order_handler.go
│   └── client/
│       └── inventory_client.go
└── migrations/
└── 001_init.up.sql