package api

import (
	"context"
	"encoding/json"
	"fmt"
	"main/models"
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
		//get all path params from url
		params := mux.Vars(r)
		model := params["sneaker"]
		sneaker, err := s.sneakerService.GetSneakerByModel(context.TODO(), model)
		if err != nil {
			json.NewEncoder(w).Encode("No sneakers found")
		}
		json.NewEncoder(w).Encode(sneaker) //sync.Mutex from protobuf, Encode copies lock value
	}
}

//createSneakerRequest contains fields when receiving a request to create a sneaker
type createSneakerRequest struct {
	Brand         string               `json:"brand"`
	Model         string               `json:"model"`
	Sku           string               `json:"sku"`
	Photos        []string             `json:"photos"`
	SiteSizePrice models.SiteSizePrice `json:"site_size_price"`
	ReleaseDate   string               `json:"release_date"`
}

func (s SneakerAPI) CreateSneaker() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createSneakerRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			json.NewEncoder(w).Encode(fmt.Sprintf("Could not parse err: %v", err))
		}
		sneaker, err := s.sneakerService.CreateSneaker(context.TODO(), req.Model, req.Brand, req.Sku, req.Photos, req.SiteSizePrice, req.ReleaseDate)
		if err != nil {
			json.NewEncoder(w).Encode(fmt.Sprintf("Unable to create err: %v", err))
		}
		json.NewEncoder(w).Encode(sneaker)
	}
}
