package domain

type Order struct {
	UserID         int64          `json:"user_id"`
	OrderNo        int64          `json:"order_no"`
	Items          []CheckoutItem `json:"items"`
	Subtotal       int64          `json:"subtotal"`
	Total          int64          `json:"total"`
	Carrier        string         `json:"carrier"`
	ShippingMethod `json:"shipping_method"`
	State          string `json:"state"`
}
