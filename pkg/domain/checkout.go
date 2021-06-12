package domain

type CheckoutItem struct {
	Model    string `json:"model"`
	Size     string `json:"size"`
	Site     string `json:"site"`
	Quantity int64  `json:"quantity"`
}

type CheckoutRequest struct {
	AuthID         string         `json:"auth_id"`
	Items          []CheckoutItem `json:"items"`
	State          string         `json:"state"`
	ShippingMethod ShippingMethod `json:"shipping_method"`
	PaymentSource  string         `json:"payment_source"`
	Carrier        string         `json:"carrier"`
}
