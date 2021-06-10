package domain

type ShippingMethod string

const (
	ShippingMethodTwoDay   ShippingMethod = "two-day"
	ShippingMethodThreeDay ShippingMethod = "three-day"
	ShippingMethodNextDay  ShippingMethod = "next-day"
)

type Carrier struct {
	Name               string                   `json:"name"`
	ShippingMethodRate map[ShippingMethod]int64 `json:"shipping_method_rates"`
}
