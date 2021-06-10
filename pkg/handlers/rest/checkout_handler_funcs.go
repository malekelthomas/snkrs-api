package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"snkrs/pkg/domain"
	"snkrs/pkg/services"
	"snkrs/pkg/services/conversion"
)

func executeCheckout(s services.Checkout, c conversion.CheckoutConversionService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req domain.CheckoutRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			json.NewEncoder(w).Encode(fmt.Sprintf("Could not parse err: %v", err))
		}
		order, err := c.ConvertCheckoutRequestToOrder(r.Context(), req)
		if err != nil {
			json.NewEncoder(w).Encode("invalid checkout request")
		}
		_, err = s.ProcessOrder(r.Context(), *order)
		if err != nil {
			json.NewEncoder(w).Encode(fmt.Sprintf("could not process order err: %v", err))
		}
		json.NewEncoder(w).Encode(order)
	}
}
