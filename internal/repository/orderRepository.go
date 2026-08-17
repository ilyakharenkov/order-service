package repository

import "database/sql"

type OrderRepository interface {
	CreateOrder()
}

type orderRepositoryPostgres struct {
	db *sql.DB
}

func (repository *orderRepositoryPostgres) CreateOrder() {

}
