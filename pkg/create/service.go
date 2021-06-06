package create

import "context"

//Service provides sneaker creation operations
type Service interface {
	//CreateSneaker defines a model with type sneaker to be stored in repository
	CreateSneaker(ctx context.Context, model, brand, sku string, photos []string, siteSizePrice SiteSizePrice, releaseDate string) (*Sneaker, error)
}

//Repository provides sneaker creation operations with a particular db
type Repository interface {
	//CreateSneaker stores sneaker in repository
	CreateSneaker(ctx context.Context, sneaker Sneaker) (*Sneaker, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}

}

func (s service) CreateSneaker(ctx context.Context, model, brand, sku string, photos []string, siteSizePrice SiteSizePrice, releaseDate string) (*Sneaker, error) {
	sneaker := Sneaker{
		Brand:            brand,
		Model:            model,
		Sku:              sku,
		Photos:           photos,
		ReleaseDate:      releaseDate,
		SitesSizesPrices: &siteSizePrice,
	}

	return s.r.CreateSneaker(ctx, sneaker)
}
