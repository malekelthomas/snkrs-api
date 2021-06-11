package services

import (
	"context"
	"snkrs/pkg/domain"
)

type User interface {
	//Create adds a new user
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	//GetByEmail retrieves a user by email
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
}

type UserRepository interface {
	//Create stores user in repo
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	//GetByEmail retrieves a user from repo by email
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
}

type userservice struct {
	ur UserRepository
}

func NewUserService(ur UserRepository) User {
	return &userservice{ur: ur}
}

func (u userservice) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	return u.ur.CreateUser(ctx, user)
}

func (u userservice) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return u.ur.GetUserByEmail(ctx, email)
}
