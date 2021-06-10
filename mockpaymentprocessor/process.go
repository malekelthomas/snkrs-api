package mockpaymentprocessor

import (
	"context"
	"snkrs/pkg/domain"
	"snkrs/pkg/services"
)

type MockProcessor struct {
	sr services.SneakerRepository
}

func NewMockProcessor(sr services.SneakerRepository) *MockProcessor {
	return &MockProcessor{
		sr: sr,
	}
}

func (m *MockProcessor) ProcessOrder(ctx context.Context, order domain.Order) (*domain.Order, error) {
	var err error
	order.Subtotal, err = m.CalculateSubtotal(ctx, order.Items)
	if err != nil {
		return nil, err
	}

	order.Total, err = m.CalculateTotal(ctx, order.Subtotal, order.Carrier, order.ShippingMethod)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (m *MockProcessor) CalculateSubtotal(ctx context.Context, items []domain.CheckoutItem) (int64, error) {

	var subtotal int64

	for i := range items {
		item := items[i]
		sneaker, err := m.sr.GetSneakerByModel(ctx, item.Model)
		if err != nil {
			return 0, err
		}
		subtotal += sneaker.SitesSizesPrices.SitesSizesPrices[item.Site].SizesPrices[item.Size] * item.Quantity //price of item on item.Site for item.Size

	}

	return subtotal, nil
}

func (m *MockProcessor) CalculateTotal(ctx context.Context, subtotal int64, carrier domain.Carrier, shippingMethod domain.ShippingMethod) (int64, error) {
	var total int64

	//TODO: calculate total using carrier shipping method rates
	return total, nil
}
