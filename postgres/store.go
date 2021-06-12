package postgres

import (
	"errors"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // driver
)

type Store struct {
	DB *sqlx.DB
}

func NewPostgresStore(conn string) (*Store, error) {
	if conn == "" {
		return nil, errors.New("no connection string provided")
	}
	log.Println("attempting to establish connection with db")
	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		log.Println("unable to establish connection", err)
		return nil, err
	}
	fmt.Println("established connection with postgres database")
	return &Store{DB: db}, nil

}
