package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"snkrs/pkg/domain"
	"snkrs/pkg/services"

	"github.com/gorilla/mux"
)

func createSneaker(s services.Sneaker) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req domain.CreateSneakerRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			json.NewEncoder(w).Encode(fmt.Sprintf("Could not parse err: %v", err))
		}
		sneaker, err := s.CreateSneaker(r.Context(), req.Model, req.Brand, req.Sku, req.Photos, req.SiteSizePrice, req.ReleaseDate)
		if err != nil {
			json.NewEncoder(w).Encode(fmt.Sprintf("Unable to create err: %v", err))
		}
		log.Println("stored", sneaker.Model)
		json.NewEncoder(w).Encode(sneaker)
	}
}

func getAllSneakers(s services.Sneaker) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r)
		sneakers, err := s.GetAllSneakers(r.Context())
		if err != nil {
			json.NewEncoder(w).Encode("No sneakers found")
		}
		json.NewEncoder(w).Encode(sneakers)
	}
}
func getAllBrands(s services.Sneaker) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r)
		brands, err := s.GetAllBrands(r.Context())
		if err != nil {
			json.NewEncoder(w).Encode("No brands found")
		}
		json.NewEncoder(w).Encode(brands)
	}
}

func getSneakersByBrand(s services.Sneaker) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		brand := params["brand"]
		sneakers, err := s.GetSneakersByBrand(r.Context(), brand)
		if err != nil {
			json.NewEncoder(w).Encode("No sneakers found")
		}
		json.NewEncoder(w).Encode(sneakers)
	}
}

func getSneakerByModel(s services.Sneaker) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//get query param value
		model := r.FormValue("model")
		sneaker, err := s.GetSneakerByModel(r.Context(), model)
		if err != nil {
			json.NewEncoder(w).Encode("No sneakers found")
		}
		json.NewEncoder(w).Encode(sneaker)
	}
}
