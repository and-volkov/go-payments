package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(account *Account) error
	DeleteAccount(id int) error
	UpdateAccount(account *Account) error
	GetAccountById(id int) (*Account, error)
	GetAccounts() ([]*Account, error)
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	connStr := "user=andrey password=secret dbname=gopayments sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{db: db}, nil
}

func (s *PostgresStorage) Init() error {
	return s.CreateAccountTable()
}

func (s *PostgresStorage) CreateAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account (
        id SERIAL PRIMARY KEY, 
        first_name VARCHAR(50),
        last_name VARCHAR(50),
        number UUID,
        balance INT,
        created_at TIMESTAMP
    )`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStorage) CreateAccount(account *Account) error {
	query := `INSERT INTO account (
        first_name,
        last_name,
        number,
        balance,
        created_at) VALUES ($1, $2, $3, $4, $5)`

	_, err := s.db.Exec(query,
		account.FirstName,
		account.LastName,
		account.Number,
		account.Balance,
		account.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStorage) DeleteAccount(id int) error {
	return nil
}

func (s *PostgresStorage) UpdateAccount(account *Account) error {
	return nil
}

func (s *PostgresStorage) GetAccountById(id int) (*Account, error) {
	return nil, nil
}

func (s *PostgresStorage) GetAccounts() ([]*Account, error) {
	query := `SELECT id, first_name, last_name, number, balance, created_at FROM account`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	accounts := make([]*Account, 0)
	for rows.Next() {
		account := new(Account)
		if err := rows.Scan(
			&account.ID,
			&account.FirstName,
			&account.LastName,
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
