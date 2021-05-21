package store

import (
	"main/models"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type SneakerStore struct {
	DB *sqlx.DB
}

type sneaker struct {
	price int64          `db:"price"`
	brand string         `db:"brand"`
	model string         `db:"model"`
	sku   string         `db:"sku"`
	sites pq.StringArray `db:"sites"`
}

func (s SneakerStore) GetAllSneakers() ([]models.Sneaker, error) {

	//get values from db scan into store sneaker type
	var sneakers []sneaker
	if err := s.DB.Select(&sneakers, `SELECT * FROM sneakers`); err != nil {
		return nil, err
	}

	//convert and return array of model type
	var converted []models.Sneaker
	for _, s := range sneakers {
		converted = append(converted, s.ToSneaker())
	}

	return converted, nil

}

func (s sneaker) ToSneaker() models.Sneaker {

	//convert values returned from db to site_sold_on type so it's methods can be used
	var sites []models.Sneaker_SiteSoldOn
	for _, site := range s.sites {
		val := models.Sneaker_SiteSoldOn_value[site]
		sites = append(sites, models.Sneaker_SiteSoldOn(val))
	}
	return models.Sneaker{
		Price: s.price,
		Brand: s.brand,
		Model: s.model,
		Sku:   s.sku,
		Sites: sites,
	}
}
