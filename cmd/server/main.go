package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"snkrs/mockpaymentprocessor"
	"snkrs/mongo"
	"snkrs/pkg/generator"
	"snkrs/pkg/handlers/rest"
	"snkrs/pkg/services"
	"snkrs/pkg/services/conversion"
	"snkrs/postgres"
	"snkrs/sonyflake"
)

func main() {
	//configure storage type

	storageType := os.Getenv("STORAGE_FLAG")

	//init services
	var (
		sneakerService            services.Sneaker
		checkoutService           services.Checkout
		checkoutConversionService conversion.CheckoutConversionService
	)

	switch storageType {
	case "0":
		s, err := mongo.NewMongoStore(os.Getenv("MONGO_CONN"))
		if err != nil {
			panic(err)
		}
		sneakerService = services.NewSneakerService(s)
		paymentProcessor := mockpaymentprocessor.NewMockProcessor(s)
		checkoutService = services.NewCheckoutService(paymentProcessor)
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

		sneakerService = services.NewSneakerService(s)
		paymentProcessor := mockpaymentprocessor.NewMockProcessor(s)
		checkoutService = services.NewCheckoutService(paymentProcessor)
	}

	//set generator for order numbers
	noGenerator := sonyflake.NewSonyflake()
	orderNoGenerator := generator.NewOrderNumberGenerator(noGenerator)
	checkoutConversionService = conversion.NewCheckoutConversionService(orderNoGenerator)

	//pass services to handlers
	router := rest.Handler(rest.Services{
		SneakerService:            sneakerService,
		CheckoutService:           checkoutService,
		CheckoutConversionService: checkoutConversionService,
	})

	fmt.Println("listening on port 7000")
	log.Fatal(http.ListenAndServe(":7000", router))
}
