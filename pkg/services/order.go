package services

import (
	"context"
	"snkrs/pkg/domain"
)

type Order interface {
	CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error)
}

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error)
}

type orderservice struct {
	r OrderRepository
}

func NewOrderService(r OrderRepository) Order {
	return &orderservice{r: r}
}

func (o *orderservice) CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	return o.r.CreateOrder(ctx, order)
}
