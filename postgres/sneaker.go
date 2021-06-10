package postgres

import (
	"context"
	"snkrs/pkg/domain"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type sneaker struct {
	id     int64          `db:"id"`
	brand  int64          `db:"brand"`
	model  string         `db:"model"`
	photos pq.StringArray `db:"photos"`
}

func (s *Store) CreateSneaker(ctx context.Context, sneaker domain.Sneaker) (*domain.Sneaker, error) {

	if s.DB != nil {
		brandID, err := s.GetBrandIDByName(ctx, strings.ToLower(sneaker.Brand))
		if err != nil {
			return nil, err
		}
		tx, err := s.DB.Begin()
		if err != nil {
			return nil, err
		}
		//store sneaker in inventory
		var inventoryID int64
		if err := tx.QueryRow(`INSERT INTO sneaker_inventory (sku, model_name) VALUES ($1, $2) RETURNING id`, sneaker.Sku, sneaker.Model).Scan(&inventoryID); err != nil {
			return nil, err
		}

		//add sneaker to 'catalog'
		//convert string array to pq array
		var photos pq.StringArray

		photos = append(photos, sneaker.Photos...)

		var sneakerID int64
		if err := tx.QueryRow(`INSERT INTO sneakers (brand_id, model_name, photos) VALUES ($1, $2, $3) RETURNING id`, brandID, sneaker.Model, photos).Scan(&sneakerID); err != nil {
			return nil, err
		}

		//add site, size, and price info
		if err := s.StoreSiteSizePrice(ctx, tx, sneakerID, inventoryID, sneaker.SitesSizesPrices); err != nil {
			return nil, err
		}
		if err := tx.Commit(); err != nil {
			return nil, err
		}
	}

	return &sneaker, nil
}

func (s *Store) GetAllSneakers(ctx context.Context) ([]domain.Sneaker, error) {
	//get values from db scan into store sneaker type
	var sneakers []sneaker
	if err := s.DB.Select(&sneakers, `SELECT * FROM sneakers`); err != nil {
		return nil, err
	}

	//convert and return array of model type
	var converted []domain.Sneaker
	var err error
	for _, sneaker := range sneakers {
		convertedSneaker := sneaker.ToSneaker(s.DB)
		convertedSneaker.SitesSizesPrices, err = s.GetSitesSizesPrices(ctx, sneaker.id)
		if err != nil {
			return nil, err
		}
		converted = append(converted, convertedSneaker)
	}

	return converted, nil

}

func (s *Store) GetSneakersByBrandID(ctx context.Context, brandID int64) ([]domain.Sneaker, error) {
	//get values from db scan into store sneaker type
	var sneakers []sneaker
	if err := s.DB.Select(&sneakers, `SELECT * FROM sneakers WHERE brand_id=$1`, brandID); err != nil {
		return nil, err
	}

	//convert and return array of model type
	var converted []domain.Sneaker
	var err error
	for _, sneaker := range sneakers {
		convertedSneaker := sneaker.ToSneaker(s.DB)
		convertedSneaker.SitesSizesPrices, err = s.GetSitesSizesPrices(ctx, sneaker.id)
		if err != nil {
			return nil, err
		}
		converted = append(converted, convertedSneaker)
	}

	return converted, nil

}

func (s *Store) GetSneakersByBrand(ctx context.Context, brand string) ([]domain.Sneaker, error) {
	//get values from db scan into store sneaker type
	var sneakers []sneaker
	var err error

	brandID, err := s.GetBrandIDByName(ctx, brand)
	if err != nil {
		return nil, err
	}
	if err := s.DB.Select(&sneakers, `SELECT * FROM sneakers WHERE brand_id=$1`, brandID); err != nil {
		return nil, err
	}

	//convert and return array of model type
	var converted []domain.Sneaker
	for _, sneaker := range sneakers {
		convertedSneaker := sneaker.ToSneaker(s.DB)
		convertedSneaker.SitesSizesPrices, err = s.GetSitesSizesPrices(ctx, sneaker.id)
		if err != nil {
			return nil, err
		}
		converted = append(converted, convertedSneaker)
	}

	return converted, nil

}

func (s *Store) GetSneakerBySKU(ctx context.Context, sku string) (domain.Sneaker, error) {
	var sneaker sneaker

	if err := s.DB.Get(&sneaker, `SELECT * FROM sneakers WHERE sku=$1`, sku); err != nil {
		return domain.Sneaker{}, err
	}
	return sneaker.ToSneaker(s.DB), nil
}

func (s *Store) GetSneakerByModel(ctx context.Context, model string) (domain.Sneaker, error) {
	var sneaker sneaker
	var err error
	if err := s.DB.Get(&sneaker, `SELECT * FROM sneakers WHERE model=$1`, model); err != nil {
		return domain.Sneaker{}, err
	}
	convertedSneaker := sneaker.ToSneaker(s.DB)
	convertedSneaker.SitesSizesPrices, err = s.GetSitesSizesPrices(ctx, sneaker.id)
	if err != nil {
		return domain.Sneaker{}, err
	}

	return convertedSneaker, nil
}

func (s sneaker) ToSneaker(db *sqlx.DB) domain.Sneaker {

	//convert values returned from db to site_sold_on type so it's methods can be used
	var brand string
	if err := db.Get(&brand, `SELECT name FROM brands WHERE id=$1`, s.brand); err != nil {
		return domain.Sneaker{}
	}
	return domain.Sneaker{
		Brand: brand,
		Model: s.model,
	}
}
