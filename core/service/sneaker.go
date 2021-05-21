package service

import (
	"main/models"

	"main/store"
)

type SneakerService struct {
	sneakerStore *store.SneakerStore
}

func (s SneakerService) GetAllSneakers() ([]models.Sneaker, error) {
	return s.sneakerStore.GetAllSneakers()
}
