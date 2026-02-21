package postgres

import (
	"database/sql"
	"Hotel_reservation/internal/user"
	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresDB(db *sql.DB) *PostgresStore{
	return &PostgresStore{db:db}
}

func ( p *PostgresStore) UserExsist(email string) (bool, error) {
	var exists bool

	query := `SELECT EXISTS (SELECT 1 FROM users WHERE email=$1)`

	err := p.db.QueryRow(query , email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (p *PostgresStore) UserSave(u user.User) error {	
	query := `
	INSERT INTO users (email, password)
	VALUES($1, $2)
	`
	_, err := p.db.Exec(query, u.Email, u.Hash)
	return err
}

func (p *PostgresStore) GetUserByEmail (email string) (user.User, error){
	var u user.User

	query := `
	SELECT id, email, hash
	FROM users
	WHERE email = $1
	`

	err := p.db.QueryRow(query, email).Scan(
		&u.ID,
		&u.Email,
		&u.Hash,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return user.User{},err
		}
		return user.User{}, err
	}
	return u, nil
}



func InitPostgres(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db , err
}

