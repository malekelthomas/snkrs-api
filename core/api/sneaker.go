package api

import (
	"encoding/json"
	"main/service"
	"net/http"

	"github.com/gorilla/mux"
)

type SneakerAPI struct {
	sneakerService *service.SneakerService
}

func NewSneakerAPI(sneakerService service.SneakerService) *SneakerAPI {
	return &SneakerAPI{
		sneakerService: &sneakerService,
	}
}

func (s SneakerAPI) GetAllSneakers() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sneakers, err := s.sneakerService.GetAllSneakers()
		if err != nil {
			json.NewEncoder(w).Encode("No sneakers found")
		}
		json.NewEncoder(w).Encode(sneakers)
	}
}

func (s SneakerAPI) GetSneaker() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		name := params["sneaker"]
		sneaker, err := s.sneakerService.GetSneakerByName(name)
		if err != nil {
			json.NewEncoder(w).Encode("No sneakers found")
		}
		json.NewEncoder(w).Encode(sneaker)
	}
}
