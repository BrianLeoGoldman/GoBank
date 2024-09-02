package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Storage interface {
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
	CreateAccount(*Account) (*Account, error)
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

func (s *PostgreSQLStorage) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("SELECT * FROM account")
	if err != nil {
		return nil, err
	}
	accounts := []*Account{}
	for rows.Next() {
		account := new(Account)
		if err := rows.Scan(
			&account.ID,
			&account.Firstname,
			&account.Lastname,
			&account.Number,
			&account.Balance,
			&account.CreatedAt,
		); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (s *PostgreSQLStorage) GetAccountByID(int) (*Account, error) {
	return nil, nil
}

func (s *PostgreSQLStorage) CreateAccount(account *Account) (*Account, error) {
	query := `INSERT INTO account
        (first_name, last_name, number, balance, created_at)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, first_name, last_name, number, balance, created_at`

	// Crear un objeto Account para almacenar la cuenta creada
	createdAccount := &Account{}

	// Ejecutar la consulta y escanear el resultado
	err := s.db.QueryRow(query,
		account.Firstname,
		account.Lastname,
		account.Number,
		account.Balance,
		account.CreatedAt,
	).Scan(
		&createdAccount.ID,
		&createdAccount.Firstname,
		&createdAccount.Lastname,
		&createdAccount.Number,
		&createdAccount.Balance,
		&createdAccount.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return createdAccount, nil
}

func (s *PostgreSQLStorage) DeleteAccount(int) error {
	return nil
}

func (s *PostgreSQLStorage) UpdateAccount(*Account) error {
	return nil
}
