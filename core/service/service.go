package service

import (
	"context"
	"main/store"
)

type Services struct {
	SneakerService *SneakerService
}

var AllServices Services

func Init(ctx context.Context) {
	store.Init(ctx)
	AllServices.SneakerService = NewSneakerService(store.NewSneakerStore(&store.Conn))
}
