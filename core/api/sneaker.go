package api

import (
	"context"
	"encoding/json"
	"fmt"
	"main/service"
	"net/http"

	"github.com/gorilla/mux"
)

//SneakerAPI handles http requests that require the sneaker service
type SneakerAPI struct {
	sneakerService *service.SneakerService
}

func NewSneakerAPI(sneakerService service.SneakerService) *SneakerAPI {
	return &SneakerAPI{
		sneakerService: &sneakerService,
	}
}

//return a closure so gorilla mux will pass in args to func returned

func (s SneakerAPI) GetAllSneakers() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sneakers, err := s.sneakerService.GetAllSneakers(context.TODO())
		if err != nil {
			json.NewEncoder(w).Encode("No sneakers found")
		}
		json.NewEncoder(w).Encode(sneakers)
	}
}

func (s SneakerAPI) GetSneaker() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		model := params["sneaker"]
		sneaker, err := s.sneakerService.GetSneakerByModel(context.TODO(), model)
		if err != nil {
			json.NewEncoder(w).Encode("No sneakers found")
		}
		json.NewEncoder(w).Encode(sneaker)
	}
}

type CreateSneakerRequest struct {
	Price  int64    `json:"price"`
	Brand  string   `json:"brand"`
	Model  string   `json:"model"`
	Sku    string   `json:"sku"`
	Sites  []string `json:"sites"`
	Photos []string `json:"photos"`
}

func (s SneakerAPI) CreateSneaker() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateSneakerRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			json.NewEncoder(w).Encode(fmt.Sprintf("Could not parse err: %v", err))
		}
		sneaker, err := s.sneakerService.CreateSneaker(context.TODO(), req.Model, req.Brand, req.Sku, req.Sites, req.Photos, req.Price)
		if err != nil {
			json.NewEncoder(w).Encode(fmt.Sprintf("Unable to create err: %v", err))
		}
		json.NewEncoder(w).Encode(sneaker)
	}
}
