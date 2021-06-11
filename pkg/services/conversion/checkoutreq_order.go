package conversion

import (
	"context"
	"snkrs/pkg/domain"
	"snkrs/pkg/generator"
)

//CheckoutService provides checkout processing operations
type CheckoutConversionService interface {
	//ConvertCheckoutRequestToOrder converts a checkout request to an order
	ConvertCheckoutRequestToOrder(ctx context.Context, req domain.CheckoutRequest) (*domain.Order, error)
}

type checkoutConversionService struct {
	o generator.OrderNumberGenerator
}

func NewCheckoutConversionService(o generator.OrderNumberGenerator) CheckoutConversionService {
	return &checkoutConversionService{o}
}

func (c checkoutConversionService) ConvertCheckoutRequestToOrder(ctx context.Context, req domain.CheckoutRequest) (*domain.Order, error) {
	var o domain.Order
	//set userID
	o.UserID = req.UserID
	//create orderNO
	var err error
	o.OrderNo, err = c.o.New()
	if err != nil {
		return nil, err
	}
	//set items
	o.Items = req.Items

	//set carrier
	o.Carrier = req.Carrier

	//set state
	o.State = req.State

	//set shipping method
	o.ShippingMethod = req.ShippingMethod

	return &o, nil
}
