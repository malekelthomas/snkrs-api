package api

import (
	"context"
	"log"
	"main/service"
	"net/http"

	"github.com/gorilla/mux"
)

func Init() {
	//init services
	service.Init(context.TODO())

	//create API
	sneakerAPI := NewSneakerAPI(*service.AllServices.SneakerService)

	r := mux.NewRouter()

	//register routes and handler funcs
	r.HandleFunc("/sneakers/", sneakerAPI.GetAllSneakers()).Methods("GET")
	r.HandleFunc("/sneakers/", sneakerAPI.CreateSneaker()).Methods("POST")
	r.HandleFunc("/sneakers/{sneaker}/", sneakerAPI.GetSneaker()).Methods("GET")
	http.Handle("/", r)

	//serve
	log.Fatal(http.ListenAndServe(":8080", r))
}
