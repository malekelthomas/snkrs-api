package domain

type CheckoutItem struct {
	Model    string `json:"model"`
	Size     string `json:"size"`
	Site     string `json:"site"`
	Quantity int64  `json:"quantity"`
	Price    int64  `json:"price"`
	Photo    string `json:"photo"`
}

type CheckoutRequest struct {
	AuthID         string         `json:"auth_id"`
	Items          []CheckoutItem `json:"items"`
	State          string         `json:"state"`
	ShippingMethod string         `json:"shipping_method"`
	PaymentSource  string         `json:"payment_source"`
	Carrier        string         `json:"carrier"`
}
