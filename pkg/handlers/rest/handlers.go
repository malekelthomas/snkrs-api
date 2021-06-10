package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"snkrs/pkg/create"
	"snkrs/pkg/domain"
	"snkrs/pkg/get"

	"github.com/gorilla/mux"
)

func Handler(s Services) *mux.Router {

	r := mux.NewRouter()

	//register routes and handler funcs
	r.HandleFunc("/sneakers/", getAllSneakers(s.Get)).Methods("GET")
	r.HandleFunc("/sneakers/", createSneaker(s.Create)).Methods("POST")
	r.HandleFunc("/sneakers", getSneakerByModel(s.Get)).Methods("GET").Queries("model", "{model}")
	r.HandleFunc("/sneakers/brands/{brand}/", getSneakersByBrand(s.Get)).Methods("GET")
	r.HandleFunc("/sneakers/brands/", getAllBrands(s.Get)).Methods("GET")
	http.Handle("/", r)

	return r
}

//createSneakerRequest contains fields when receiving a request to create a sneaker
type createSneakerRequest struct {
	Brand         string               `json:"brand"`
	Model         string               `json:"model"`
	Sku           string               `json:"sku"`
	Photos        []string             `json:"photos"`
	SiteSizePrice domain.SiteSizePrice `json:"site_size_price"`
	ReleaseDate   string               `json:"release_date"`
}

func createSneaker(c create.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createSneakerRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			json.NewEncoder(w).Encode(fmt.Sprintf("Could not parse err: %v", err))
		}
		sneaker, err := c.CreateSneaker(context.TODO(), req.Model, req.Brand, req.Sku, req.Photos, req.SiteSizePrice, req.ReleaseDate)
		if err != nil {
			json.NewEncoder(w).Encode(fmt.Sprintf("Unable to create err: %v", err))
		}
		log.Println("stored", sneaker.Model)
		json.NewEncoder(w).Encode(sneaker)
	}
}

func getAllSneakers(g get.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r)
		sneakers, err := g.GetAllSneakers(context.TODO())
		if err != nil {
			json.NewEncoder(w).Encode("No sneakers found")
		}
		json.NewEncoder(w).Encode(sneakers)
	}
}
func getAllBrands(g get.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r)
		brands, err := g.GetAllBrands(context.TODO())
		if err != nil {
			json.NewEncoder(w).Encode("No brands found")
		}
		json.NewEncoder(w).Encode(brands)
	}
}

func getSneakersByBrand(g get.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		brand := params["brand"]
		sneakers, err := g.GetSneakersByBrand(context.TODO(), brand)
		if err != nil {
			json.NewEncoder(w).Encode("No sneakers found")
		}
		json.NewEncoder(w).Encode(sneakers)
	}
}

func getSneakerByModel(g get.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//get query param value
		model := r.FormValue("model")
		sneaker, err := g.GetSneakerByModel(context.TODO(), model)
		if err != nil {
			json.NewEncoder(w).Encode("No sneakers found")
		}
		json.NewEncoder(w).Encode(sneaker)
	}
}
