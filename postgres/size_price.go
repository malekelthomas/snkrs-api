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
	ID          int64  `db:"id"`
	SiteID      int64  `db:"site_id"`
	SneakerID   int64  `db:"sneaker_id"`
	Size        string `db:"size"`
	Price       int64  `db:"price"`
	InventoryID int64  `db:"inventory_id"`
}

func (s *Store) GetSitesSizesPrices(ctx context.Context, sneakerID int64) (*domain.SiteSizePrice, error) {
	var siteSizePrice domain.SiteSizePrice
	sitesSizesPrice := make(map[string]*domain.SizePrice)
	siteSizePrice.SitesSizesPrices = sitesSizesPrice

	//get all rows for this sneaker
	rows, err := s.DB.Queryx(`SELECT * FROM site_size_price WHERE sneaker_id=$1`, sneakerID)
	if err != nil {
		return nil, err
	}

	//build out maps
	for rows.Next() {
		var row siteSizePriceRow
		if err := rows.StructScan(&row); err != nil {
			return nil, err
		}

		//get site name
		var site string
		if err := s.DB.Get(&site, `SELECT name FROM sites WHERE id=$1`, row.SiteID); err != nil {
			return nil, err
		}
		//initialize sizes and prices
		var sizePrice domain.SizePrice
		sizePriceMap := make(map[string]int64)
		sizePrice.SizesPrices = sizePriceMap

		//check if size_price map already exists for site
		if val, ok := siteSizePrice.SitesSizesPrices[site]; !ok {
			//if not set it
			siteSizePrice.SitesSizesPrices[site] = &sizePrice
		} else if ok {
			//if it does set pointer to map
			sizePrice = *val
		}
		//add size price
		sizePrice.SizesPrices[row.Size] = row.Price

	}

	return &siteSizePrice, nil
}
