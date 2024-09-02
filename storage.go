package main

import (
	"database/sql"
	"fmt"
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
	connStr := "user=postgres dbname=postgres password=pass12345 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgreSQLStorage{
		db: db,
	}, nil
}

func (s *PostgreSQLStorage) Init() error {
	return s.createAccountTable()
}

func (s *PostgreSQLStorage) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account (
		id serial PRIMARY KEY,
		first_name VARCHAR(50),
		last_name VARCHAR(50),
		number serial,
		balance serial,
		created_at TIMESTAMP
	)`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgreSQLStorage) GetAccountByID(int) (*Account, error) {
	return nil, nil
}

func (s *PostgreSQLStorage) CreateAccount(account *Account) error {
	query := `INSERT INTO account
		(first_name, last_name, number, balance, created_at)
		values
		($1, $2, $3, $4, $5)`
	resp, err := s.db.Query(query,
		account.Firstname,
		account.Lastname,
		account.Number,
		account.Balance,
		account.CreatedAt,
	)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", resp)
}

func (s *PostgreSQLStorage) DeleteAccount(int) error {
	return nil
}

func (s *PostgreSQLStorage) UpdateAccount(*Account) error {
	return nil
}
