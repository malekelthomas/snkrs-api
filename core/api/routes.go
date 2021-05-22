package api

import (
	"context"
	"log"
	"main/service"
	"net/http"

	"github.com/gorilla/mux"
)

func Init() {
	service.Init(context.TODO())
	sneakerAPI := NewSneakerAPI(*service.AllServices.SneakerService)

	r := mux.NewRouter()
	r.HandleFunc("/sneakers/", sneakerAPI.GetAllSneakers()).Methods("GET")
	r.HandleFunc("/sneakers/", sneakerAPI.CreateSneaker()).Methods("POST")
	r.HandleFunc("/sneakers/{sneaker}/", sneakerAPI.GetSneaker()).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
