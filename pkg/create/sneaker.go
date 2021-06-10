package create

import (
	"context"
	"snkrs/pkg/domain"
)

//Service provides sneaker creation operations
type SneakerService interface {
	//CreateSneaker defines a model with type sneaker to be stored in repository
	CreateSneaker(ctx context.Context, model, brand, sku string, photos []string, siteSizePrice domain.SiteSizePrice, releaseDate string) (*domain.Sneaker, error)
}

//Repository provides sneaker creation operations with a particular db
type SneakerRepository interface {
	//CreateSneaker stores sneaker in repository
	CreateSneaker(ctx context.Context, sneaker domain.Sneaker) (*domain.Sneaker, error)
}

type sneakerservice struct {
	r SneakerRepository
}

func NewSneakerService(r SneakerRepository) SneakerService {
	return &sneakerservice{r}

}

func (s sneakerservice) CreateSneaker(ctx context.Context, model, brand, sku string, photos []string, siteSizePrice domain.SiteSizePrice, releaseDate string) (*domain.Sneaker, error) {
	sneaker := domain.Sneaker{
		Brand:            brand,
		Model:            model,
		Sku:              sku,
		Photos:           photos,
		ReleaseDate:      releaseDate,
		SitesSizesPrices: &siteSizePrice,
	}

	return s.r.CreateSneaker(ctx, sneaker)
}
