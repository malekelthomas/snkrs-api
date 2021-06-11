package postgres

import (
	"context"
	"snkrs/pkg/domain"
)

func (s *Store) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	var userID int64

	//get user role id
	var uRoleID int64
	if err := s.DB.Get(&uRoleID, `SELECT id FROM user_roles WHERE user_role=$1`, user.UserRole); err != nil {
		return nil, err
	}

	if err := s.DB.QueryRow(`INSERT INTO users (
			first_name,
			last_name,
			email,
			user_role_id,
			auth_id
		)
		VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		user.FirstName,
		user.LastName,
		user.Email,
		uRoleID,
		user.AuthID,
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
