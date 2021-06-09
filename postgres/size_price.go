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
