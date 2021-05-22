package api

import (
	"main/service"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	service.Init()
	sneakerAPI := NewSneakerAPI(*service.AllServices.SneakerService)
	r := mux.NewRouter()
	r.HandleFunc("/sneakers", sneakerAPI.GetAllSneakers())
	r.HandleFunc("sneakers/{sneaker}/", sneakerAPI.GetSneaker())
	http.Handle("/", r)
}
