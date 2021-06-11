package postgres

import (
	"context"
	"snkrs/pkg/domain"
)

func (s *Store) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	var userID int64
	if err := s.DB.QueryRow(`INSERT INTO users (
			first_name,
			last_name,
			email,
			password,
			user_role_id,
			auth_id
		)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		user.FirstName,
		user.LastName,
		user.Email,
	).Scan(&userID); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Store) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	if err := s.DB.Get(&user, `SELECT * FROM users WHERE email=$1`, email); err != nil {
		return nil, err
	}

	return &user, nil
}
