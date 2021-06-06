package store

import (
	"context"
	"errors"
	"log"

	"snkrs/pkg/create"
	"snkrs/pkg/get"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Store struct {
	DB *sqlx.DB
}

type sneaker struct {
	price int64          `db:"price"`
	brand string         `db:"brand"`
	model string         `db:"model"`
	sku   string         `db:"sku"`
	sites pq.StringArray `db:"sites"`
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

func (s *Store) CreateSneaker(ctx context.Context, sneaker create.Sneaker) (*create.Sneaker, error) {

	if s.DB != nil {
		//
	}

	return &sneaker, nil
}

func (s *Store) GetAllSneakers(ctx context.Context) ([]get.Sneaker, error) {
	//get values from db scan into store sneaker type
	var sneakers []sneaker
	if err := s.DB.Select(&sneakers, `SELECT * FROM sneakers`); err != nil {
		return nil, err
	}

	//convert and return array of model type
	var converted []get.Sneaker
	for _, s := range sneakers {
		converted = append(converted, s.ToSneaker())
	}

	return converted, nil

}

func (s *Store) GetSneakerByModel(ctx context.Context, model string) (get.Sneaker, error) {
	var sneaker sneaker

	if err := s.DB.Get(&sneaker, `SELECT * FROM sneakers WHERE model=$1`, model); err != nil {
		return get.Sneaker{}, err
	}
	return sneaker.ToSneaker(), nil
}

func (s sneaker) ToSneaker() get.Sneaker {

	//convert values returned from db to site_sold_on type so it's methods can be used
	return get.Sneaker{
		Brand: s.brand,
		Model: s.model,
		Sku:   s.sku,
	}
}
