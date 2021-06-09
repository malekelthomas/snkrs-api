package postgres

import (
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	DB *sqlx.DB
}

func NewPostgresStore(conn string) (*Store, error) {
	if conn == "" {
		return nil, errors.New("no connection string provided")
	}
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		log.Println("unable to establish connection", err)
		return nil, err
	}
	return &Store{DB: db}, nil

}
