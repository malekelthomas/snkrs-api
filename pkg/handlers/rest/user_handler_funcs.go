package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"snkrs/pkg/domain"
	"snkrs/pkg/services"
)

func registerUser(u services.User) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req domain.User
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			json.NewEncoder(w).Encode(fmt.Sprintf("Could not parse err: %v", err))
		}
		if req.AuthID == "" {
			json.NewEncoder(w).Encode(fmt.Sprintf("invalid user: %v", req))
		} else {
			user, err := u.CreateUser(r.Context(), &req)
			if err != nil {
				json.NewEncoder(w).Encode(fmt.Sprintf("invalid user err: %v", err))
			}
			json.NewEncoder(w).Encode(user)
		}
	}
}
