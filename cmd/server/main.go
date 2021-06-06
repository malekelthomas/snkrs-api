package main

import (
	"log"
	"net/http"
	"os"
	"snkrs/mongo"
	"snkrs/pkg/create"
	"snkrs/pkg/get"
	"snkrs/pkg/http/rest"
)

func main() {
	//configure storage type

	storageType := os.Getenv("STORAGE_FLAG")

	//init services
	var (
		getService    get.Service
		createService create.Service
	)

	switch storageType {
	case "0":
		s, err := mongo.NewMongoStore(os.Getenv("MONGO_CONN"))
		if err != nil {
			panic(err)
		}

		getService = get.NewService(s)
		createService = create.NewService(s)

	}

	router := rest.Handler(rest.Services{Get: getService, Create: createService})
	log.Fatal(http.ListenAndServe(":8080", router))
}
