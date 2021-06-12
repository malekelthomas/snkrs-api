package postgres

import (
	"context"
	"errors"
	"snkrs/pkg/domain"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type sneaker struct {
	ID          int64          `db:"id"`
	Brand       int64          `db:"brand_id"`
	Model       string         `db:"model_name"`
	ReleaseDate *time.Time     `db:"release_date"`
	Photos      pq.StringArray `db:"photos"`
}

var layout = "01/02/2006"

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

		//check if release date provided
		var releaseDate time.Time
		if sneaker.ReleaseDate != "" {
			releaseDate, err = time.Parse(layout, sneaker.ReleaseDate)
			if err != nil {
				return nil, err
			}

		}

		//store sneaker in inventory
		var inventoryID int64
		if err := tx.QueryRow(`INSERT INTO sneaker_inventory (sku, model_name) VALUES ($1, $2) RETURNING id`, sneaker.Sku, sneaker.Model).Scan(&inventoryID); err != nil {
			return nil, err
		}

		//add sneaker to 'catalog'
		//convert string array to pq array

		photos := pq.StringArray(sneaker.Photos)

		var sneakerID int64
		if err := tx.QueryRow(`INSERT INTO sneakers (brand_id, model_name, photos, release_date) VALUES ($1, $2, $3, $4) RETURNING id`, brandID, sneaker.Model, photos, releaseDate).Scan(&sneakerID); err != nil {
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
	var err error
	rows, err := s.DB.Queryx(`SELECT * FROM sneakers`)
	if err != nil {
		return nil, err
	}

	//convert and return array of model type
	var converted []domain.Sneaker

	for rows.Next() {
		var sneaker sneaker
		if err := rows.StructScan(&sneaker); err != nil {
			return nil, err
		}
		convertedSneaker := sneaker.ToSneaker(s.DB)
		if convertedSneaker == nil {
			return nil, errors.New("unable to convert from table struct to model")
		}
		/* convertedSneaker.SitesSizesPrices, err = s.GetSitesSizesPrices(ctx, sneaker.ID)
		if err != nil {
			return nil, err
		} */
		converted = append(converted, *convertedSneaker)
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
	for _, sneaker := range sneakers {
		convertedSneaker := sneaker.ToSneaker(s.DB)
		/* convertedSneaker.SitesSizesPrices, err = s.GetSitesSizesPrices(ctx, sneaker.ID)
		if err != nil {
			return nil, err
		} */
		converted = append(converted, *convertedSneaker)
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
		/* convertedSneaker.SitesSizesPrices, err = s.GetSitesSizesPrices(ctx, sneaker.ID)
		if err != nil {
			return nil, err
		} */
		converted = append(converted, *convertedSneaker)
	}

	return converted, nil

}

func (s *Store) GetSneakerBySKU(ctx context.Context, sku string) (domain.Sneaker, error) {
	var sneaker sneaker

	if err := s.DB.Get(&sneaker, `SELECT * FROM sneakers WHERE sku=$1`, sku); err != nil {
		return domain.Sneaker{}, err
	}
	return *sneaker.ToSneaker(s.DB), nil
}

func (s *Store) GetSneakerByModel(ctx context.Context, model string) (domain.Sneaker, error) {
	var sneaker sneaker
	var err error
	if err := s.DB.Get(&sneaker, `SELECT * FROM sneakers WHERE model_name=$1`, model); err != nil {
		return domain.Sneaker{}, err
	}
	convertedSneaker := sneaker.ToSneaker(s.DB)
	convertedSneaker.SitesSizesPrices, err = s.GetSitesSizesPrices(ctx, sneaker.ID)
	if err != nil {
		return domain.Sneaker{}, err
	}

	return *convertedSneaker, nil
}

func (s *Store) GetSneakerIDByModel(ctx context.Context, model string) (int64, error) {
	var sneaker sneaker
	if err := s.DB.Get(&sneaker, `SELECT id FROM sneakers WHERE model_name=$1`, model); err != nil {
		return 0, err
	}
	return sneaker.ID, nil
}

func (s *Store) GetSneakerInventoryIDByModel(ctx context.Context, model string) (int64, error) {
	var sneaker sneaker
	if err := s.DB.Get(&sneaker, `SELECT id FROM sneaker_inventory WHERE model_name=$1`, model); err != nil {
		return 0, err
	}
	return sneaker.ID, nil
}
func (s *Store) GetSneakerInventoryIDBySneakerID(ctx context.Context, sneakerID int64) (int64, error) {
	var sneaker sneaker
	if err := s.DB.Get(&sneaker, `SELECT id FROM sneaker_inventory WHERE sneaker_id=$1`, sneakerID); err != nil {
		return 0, err
	}
	return sneaker.ID, nil
}

func (s sneaker) ToSneaker(db *sqlx.DB) *domain.Sneaker {

	//convert values returned from db to site_sold_on type so it's methods can be used
	var brand string
	if err := db.Get(&brand, `SELECT name FROM brands WHERE id=$1`, s.Brand); err != nil {
		return nil
	}

	return &domain.Sneaker{
		Brand:       brand,
		Model:       s.Model,
		Photos:      s.Photos,
		ReleaseDate: s.ReleaseDate.Format(layout),
	}
}
