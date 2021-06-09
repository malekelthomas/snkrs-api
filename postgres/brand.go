package postgres

import "context"

func (s *Store) GetAllBrands(ctx context.Context) ([]string, error) {
	//get values from db scan into store sneaker type
	var brands []string
	if err := s.DB.Select(&brands, `SELECT name FROM brands`); err != nil {
		return nil, err
	}

	return brands, nil

}

func (s *Store) GetBrandByID(ctx context.Context, id int64) (string, error) {
	var brand string
	if err := s.DB.Get(&brand, `SELECT name FROM brands WHERE id=$1`, id); err != nil {
		return "", err
	}
	return brand, nil
}

func (s *Store) GetBrandIDByName(ctx context.Context, name string) (int64, error) {
	var brand int64
	if err := s.DB.Get(&brand, `SELECT id FROM brands WHERE name=$1`, name); err != nil {
		return 0, err
	}
	return brand, nil
}
