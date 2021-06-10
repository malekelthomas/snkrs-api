package postgres

import (
	"context"
	"database/sql"
	"snkrs/pkg/domain"
)

func (s *Store) StoreSiteSizePrice(ctx context.Context, tx *sql.Tx, sneakerID, inventoryID int64, info *domain.SiteSizePrice) error {

	if tx == nil {
		var err error
		tx, err = s.DB.Begin()
		if err != nil {
			return err
		}
	}
	for site := range info.SitesSizesPrices {
		var siteID int64
		if err := s.DB.Get(&siteID, `SELECT id FROM sites WHERE name=$1`, site); err != nil {
			return err
		}
		sizes_prices := info.SitesSizesPrices[site]
		for size := range sizes_prices.SizesPrices {
			price := sizes_prices.SizesPrices[size]
			if _, err := tx.Exec(`INSERT INTO site_size_price (site_id, sneaker_id, size, price, inventory_id) VALUES ($1, $2, $3, $4, $5)`, siteID, sneakerID, size, price, inventoryID); err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return nil
}

type siteSizePriceRow struct {
	id         int64  `db:"id"`
	site_id    int64  `db:"site_id"`
	sneaker_id int64  `db:"sneaker_id"`
	size       string `db:"size"`
	price      int64  `db:"price"`
}

func (s *Store) GetSitesSizesPrices(ctx context.Context, sneakerID int64) (*domain.SiteSizePrice, error) {
	var info *domain.SiteSizePrice
	var rows []siteSizePriceRow

	//get all rows for this sneaker
	if err := s.DB.Select(&rows, `SELECT * FROM site_size_price WHERE sneaker_id=$1`, sneakerID); err != nil {
		return nil, err
	}

	//build out maps
	for i := range rows {
		row := rows[i]

		//get site name
		var site string
		if err := s.DB.Get(&site, `SELECT name FROM sites WHERE id=$1`, row.site_id); err != nil {
			return nil, err
		}
		//initialize sizes and prices
		var sizePrice *domain.SizePrice

		//check if size_price map already exists for site
		if val, ok := info.SitesSizesPrices[site]; !ok {
			//if not set it
			info.SitesSizesPrices[site] = sizePrice
		} else if ok {
			//if it does set pointer to map
			sizePrice = val
		}
		//add size price
		sizePrice.SizesPrices[row.size] = row.price

	}

	return info, nil
}
