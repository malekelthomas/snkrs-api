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
	//GetSneakerByBrandWPagination returns sneakers with given brand paginated
	GetSneakersByBrandWPagination(ctx context.Context, brand string, limit, offset int64) ([]domain.Sneaker, error)
	//GetAllSneakers returns all sneakers
	GetAllSneakers(ctx context.Context) ([]domain.Sneaker, error)
	//GetAllSneakersWPagination returns all sneakers paginated
	GetAllSneakersWPagination(ctx context.Context, limit, offset int64) ([]domain.Sneaker, error)
	//GetBrands returns all brands
	GetAllBrands(ctx context.Context) ([]string, error)
	//GetSneakerCount returns number of sneakers
	GetSneakerCount(ctx context.Context) (int64, error)
	//GetSneakerCount returns number of sneakers for a brand
	GetSneakerCountByBrand(ctx context.Context, brand string) (int64, error)
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
	//GetSneakerByBrandWPagination returns sneakers with given brand paginated
	GetSneakersByBrandWPagination(ctx context.Context, brand string, limit, offset int64) ([]domain.Sneaker, error)
	//GetAllSneakers returns all sneakers
	GetAllSneakers(ctx context.Context) ([]domain.Sneaker, error)
	//GetAllSneakersWPagination returns all sneakers paginated
	GetAllSneakersWPagination(ctx context.Context, limit, offset int64) ([]domain.Sneaker, error)
	//GetBrands returns all brands
	GetAllBrands(ctx context.Context) ([]string, error)
	//GetSneakerCount returns number of sneakers
	GetSneakerCount(ctx context.Context) (int64, error)
	//GetSneakerCount returns number of sneakers for a brand
	GetSneakerCountByBrand(ctx context.Context, brand string) (int64, error)
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

//GetSneakerByBrandWPagination returns sneakers with given brand paginated
func (s sneakerservice) GetSneakersByBrandWPagination(ctx context.Context, brand string, limit, offset int64) ([]domain.Sneaker, error) {
	return s.r.GetSneakersByBrandWPagination(ctx, brand, limit, offset)
}

func (s sneakerservice) GetAllSneakers(ctx context.Context) ([]domain.Sneaker, error) {
	return s.r.GetAllSneakers(ctx)
}

func (s sneakerservice) GetAllSneakersWPagination(ctx context.Context, limit, offset int64) ([]domain.Sneaker, error) {
	return s.r.GetAllSneakersWPagination(ctx, limit, offset)
}
func (s sneakerservice) GetAllBrands(ctx context.Context) ([]string, error) {
	return s.r.GetAllBrands(ctx)
}

func (s sneakerservice) GetSneakerCount(ctx context.Context) (int64, error) {
	return s.r.GetSneakerCount(ctx)
}

func (s sneakerservice) GetSneakerCountByBrand(ctx context.Context, brand string) (int64, error) {
	return s.r.GetSneakerCountByBrand(ctx, brand)
}
