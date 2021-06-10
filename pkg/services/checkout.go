package services

import (
	"context"
	"snkrs/pkg/domain"
)

//CheckoutService provides checkout processing operations
type Checkout interface {
	//Process processes a new order
	ProcessOrder(ctx context.Context, order domain.Order) (*domain.Order, error)
}

type PaymentProcessor interface {
	//Process processes a new order using the payment processor's methods
	ProcessOrder(ctx context.Context, order domain.Order) (*domain.Order, error)
}

type checkout struct {
	p PaymentProcessor
}

func NewCheckoutService(p PaymentProcessor) Checkout {
	return &checkout{p}
}

func (c checkout) ProcessOrder(ctx context.Context, order domain.Order) (*domain.Order, error) {
	return c.p.ProcessOrder(ctx, order)
}
