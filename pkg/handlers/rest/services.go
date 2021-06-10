package rest

import (
	"snkrs/pkg/services"
	"snkrs/pkg/services/conversion"
)

type Services struct {
	SneakerService            services.Sneaker
	CheckoutConversionService conversion.CheckoutConversionService
	CheckoutService           services.Checkout
}
