package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Handler(s Services) *mux.Router {

	r := mux.NewRouter()

	//register routes and handler funcs

	//SNEAKER ROUTES
	r.HandleFunc("/sneakers/", getAllSneakers(s.SneakerService)).Methods("GET")
	r.HandleFunc("/sneakers/", createSneaker(s.SneakerService)).Methods("POST")
	r.HandleFunc("/sneakers", getSneakerByModel(s.SneakerService)).Methods("GET").Queries("model", "{model}")
	r.HandleFunc("/sneakers/brands/{brand}/", getSneakersByBrand(s.SneakerService)).Methods("GET")
	r.HandleFunc("/sneakers/brands/", getAllBrands(s.SneakerService)).Methods("GET")

	//CHECKOUT ROUTES
	r.HandleFunc("/checkout/execute/", executeCheckout(s.CheckoutService, s.CheckoutConversionService)).Methods("POST")

	//USER ROUTES
	r.HandleFunc("/users/register/", registerUser(s.UserService)).Methods("POST")
	http.Handle("/", r)

	return r
}
