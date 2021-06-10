package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"snkrs/mongo"
	"snkrs/pkg/create"
	"snkrs/pkg/get"
	"snkrs/pkg/handlers/rest"
	"snkrs/postgres"
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
	case "1":
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
		)
		s, err := postgres.NewPostgresStore(dsn)
		if err != nil {
			panic(err)
		}

		getService = get.NewService(s)
		createService = create.NewService(s)

	}

	router := rest.Handler(rest.Services{Get: getService, Create: createService})
	fmt.Println("listening on port 7000")
	log.Fatal(http.ListenAndServe(":7000", router))
}
