package service

import (
	"context"
	"main/models"

	"main/store"
)

type SneakerService struct {
	sneakerStore *store.SneakerStore
}

func NewSneakerService(sneakerStore *store.SneakerStore) *SneakerService {
	return &SneakerService{
		sneakerStore: sneakerStore,
	}
}
func (s SneakerService) GetAllSneakers(ctx context.Context) ([]models.Sneaker, error) {
	return s.sneakerStore.GetAllSneakers(ctx)
}

func (s SneakerService) GetSneakerByModel(ctx context.Context, model string) (models.Sneaker, error) {
	return s.sneakerStore.GetSneakerByModel(ctx, model)
}

func (s SneakerService) CreateSneaker(ctx context.Context, model, brand, sku string, photos []string, siteSizePrice models.SiteSizePrice) (*models.Sneaker, error) {
	return s.sneakerStore.CreateSneaker(ctx, model, brand, sku, photos, siteSizePrice)
}
