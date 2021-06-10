package domain

type Sneaker struct {
	Brand            string         `json:"brand,omitempty"`
	Model            string         `json:"model,omitempty"`
	Sku              string         `json:"sku,omitempty"`
	ReleaseDate      string         `json:"release_date,omitempty"`
	Photos           []string       `json:"photos,omitempty"`
	SitesSizesPrices *SiteSizePrice `json:"sites_sizes_prices,omitempty"`
}

type SiteSizePrice struct {
	SitesSizesPrices map[string]*SizePrice `json:"sites_sizes_prices,omitempty"`
}

type SizePrice struct {
	SizesPrices map[string]int64 `json:"sizes_prices,omitempty"`
}

//CreateSneakerRequest contains fields when receiving a request to create a sneaker
type CreateSneakerRequest struct {
	Brand         string        `json:"brand"`
	Model         string        `json:"model"`
	Sku           string        `json:"sku"`
	Photos        []string      `json:"photos"`
	SiteSizePrice SiteSizePrice `json:"site_size_price"`
	ReleaseDate   string        `json:"release_date"`
}

type SiteSoldOn int32

const (
	SiteSoldOn_STOCKX       SiteSoldOn = 0
	SiteSoldOn_NIKE         SiteSoldOn = 1
	SiteSoldOn_ADIDAS       SiteSoldOn = 2
	SiteSoldOn_PUMA         SiteSoldOn = 3
	SiteSoldOn_STADIUMGOODS SiteSoldOn = 4
	SiteSoldOn_FLIGHTCLUB   SiteSoldOn = 5
)

// Enum value maps for SiteSoldOn.
var (
	SiteSoldOn_name = map[int32]string{
		0: "STOCKX",
		1: "NIKE",
		2: "ADIDAS",
		3: "PUMA",
		4: "STADIUMGOODS",
		5: "FLIGHTCLUB",
	}
	SiteSoldOn_value = map[string]int32{
		"STOCKX":       0,
		"NIKE":         1,
		"ADIDAS":       2,
		"PUMA":         3,
		"STADIUMGOODS": 4,
		"FLIGHTCLUB":   5,
	}
)
