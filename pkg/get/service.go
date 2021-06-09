package get

import (
	"context"
	"snkrs/pkg/domain"
)

//Service provides domain.Sneaker retrieval operations
type Service interface {
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

//Repository provides domain.Sneaker retrieval operations from storage
type Repository interface {
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

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s service) GetSneakerBySKU(ctx context.Context, sku string) (domain.Sneaker, error) {
	return s.r.GetSneakerBySKU(ctx, sku)
}

func (s service) GetSneakerByModel(ctx context.Context, model string) (domain.Sneaker, error) {
	return s.r.GetSneakerByModel(ctx, model)
}
func (s service) GetSneakersByBrand(ctx context.Context, brand string) ([]domain.Sneaker, error) {
	return s.r.GetSneakersByBrand(ctx, brand)
}

func (s service) GetAllSneakers(ctx context.Context) ([]domain.Sneaker, error) {
	return s.r.GetAllSneakers(ctx)
}
func (s service) GetAllBrands(ctx context.Context) ([]string, error) {
	return s.r.GetAllBrands(ctx)
}
