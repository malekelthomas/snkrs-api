package services

import (
	"context"
	"snkrs/pkg/domain"
)

//sneakerservice provides sneaker creation operations
type Sneaker interface {
	//CreateSneaker defines a model with type sneaker to be stored in repository
	CreateSneaker(ctx context.Context, model, brand, sku string, photos []string, siteSizePrice domain.SiteSizePrice, releaseDate string) (*domain.Sneaker, error)
	//GetSneakerBySKU returns domain.Sneaker with given sku
	GetSneakerBySKU(ctx context.Context, sku string) (domain.Sneaker, error)
	//GetSneakerByModel returns domain.Sneaker with given model name
	GetSneakerByModel(ctx context.Context, model string) (domain.Sneaker, error)
	//GetSneakerByBrand returns sneakers with given brand
	GetSneakersByBrand(ctx context.Context, brand string) ([]domain.Sneaker, error)
	//GetAllSneakers returns all sneakers
	GetAllSneakers(ctx context.Context) ([]domain.Sneaker, error)
	//GetBrands returns all brands
	GetAllBrands(ctx context.Context) ([]string, error)
}

//Repository provides sneaker creation operations with a particular db
type SneakerRepository interface {
	//CreateSneaker stores sneaker in repository
	CreateSneaker(ctx context.Context, sneaker domain.Sneaker) (*domain.Sneaker, error)
	//GetSneakerBySKU returns domain.Sneaker with given sku
	GetSneakerBySKU(ctx context.Context, sku string) (domain.Sneaker, error)
	//GetSneakerByModel returns domain.Sneaker with given model name
	GetSneakerByModel(ctx context.Context, model string) (domain.Sneaker, error)
	//GetSneakerByBrand returns sneakers with given brand
	GetSneakersByBrand(ctx context.Context, brand string) ([]domain.Sneaker, error)
	//GetAllSneakers returns all sneakers
	GetAllSneakers(ctx context.Context) ([]domain.Sneaker, error)
	//GetBrands returns all brands
	GetAllBrands(ctx context.Context) ([]string, error)
}

type sneakerservice struct {
	r SneakerRepository
}

func NewSneakerService(r SneakerRepository) Sneaker {
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

func (s sneakerservice) GetSneakerBySKU(ctx context.Context, sku string) (domain.Sneaker, error) {
	return s.r.GetSneakerBySKU(ctx, sku)
}

func (s sneakerservice) GetSneakerByModel(ctx context.Context, model string) (domain.Sneaker, error) {
	return s.r.GetSneakerByModel(ctx, model)
}
func (s sneakerservice) GetSneakersByBrand(ctx context.Context, brand string) ([]domain.Sneaker, error) {
	return s.r.GetSneakersByBrand(ctx, brand)
}

func (s sneakerservice) GetAllSneakers(ctx context.Context) ([]domain.Sneaker, error) {
	return s.r.GetAllSneakers(ctx)
}
func (s sneakerservice) GetAllBrands(ctx context.Context) ([]string, error) {
	return s.r.GetAllBrands(ctx)
}
