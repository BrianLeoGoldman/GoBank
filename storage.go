package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Storage interface {
	GetAccountByID(int) (*Account, error)
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
}

type PostgreSQLStorage struct {
	db *sql.DB
}

func NewPostgreSQLStorage() (*PostgreSQLStorage, error) {
	connStr := "user=postgres dbname=postgres password=gobank123 sslmode=disable"
	db, error := sql.Open("postgres", connStr)
	if error != nil {
		return nil, error
	}
	if error == db.Ping() {
		return nil, error
	}
	return &PostgreSQLStorage{
		db: db,
	}, nil
}

func (*PostgreSQLStorage) GetAccountByID(int) (*Account, error) {
	return nil, nil
}

func (*PostgreSQLStorage) CreateAccount(*Account) error {
	return nil
}

func (*PostgreSQLStorage) DeleteAccount(int) error {
	return nil
}

func (*PostgreSQLStorage) UpdateAccount(*Account) error {
	return nil
}
