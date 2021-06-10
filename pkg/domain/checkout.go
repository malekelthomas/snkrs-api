package domain

type CheckoutItem struct {
	Model    string `json:"model"`
	Quantity int64  `json:"quantity"`
}
