package postgres

import (
	"context"
	"snkrs/pkg/domain"
)

func (s *Store) CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {

	//get tax rate id
	var taxRateID int64
	if err := s.DB.Get(&taxRateID, `SELECT id FROM tax_rates WHERE state=$1`, order.State); err != nil {
		return nil, err
	}

	//get carrier id
	var carrierID int64
	if err := s.DB.Get(&carrierID, `SELECT id FROM carriers WHERE name=$1`, order.Carrier); err != nil {
		return nil, err
	}

	//get shipping method id
	var shippingMethodID int64
	if err := s.DB.Get(&shippingMethodID, `SELECT id FROM shipping_methods WHERE name=$1`, order.ShippingMethod); err != nil {
		return nil, err
	}

	//get carrier shipping method id
	var carrierShippingMethodID int64
	if err := s.DB.Get(&carrierShippingMethodID, `SELECT id FROM carrier_shipping_methods WHERE carrier_id=$1 AND shipping_method_id=$2`, carrierID, shippingMethodID); err != nil {
		return nil, err
	}

	var orderID int64
	//store order
	if err := s.DB.QueryRow(`INSERT INTO orders (
			order_no, 
			user_id, 
			subtotal, 
			tax_rate_id, 
			carrier_shipping_method_id, 
			total
		) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		order.OrderNo,
		order.UserID,
		order.Subtotal,
		taxRateID,
		carrierShippingMethodID,
		order.Total).Scan(&orderID); err != nil {
		return nil, err
	}
	//store order's items
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, err
	}

	for i := range order.Items {
		item := order.Items[i]

		sneakerID, err := s.GetSneakerIDByModel(ctx, item.Model)
		if err != nil {
			return nil, err
		}
		sneakerInventoryID, err := s.GetSneakerInventoryIDBySneakerID(ctx, sneakerID)
		if err != nil {
			return nil, err
		}
		if _, err := tx.Exec(`INSERT INTO order_items (
				order_id,
				sneaker_id,
				sneaker_inventory_id
			)
			VALUES ($1, $2, $3)`,
			orderID,
			sneakerID,
			sneakerInventoryID,
		); err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return order, nil

}
