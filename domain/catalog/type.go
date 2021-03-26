package catalog

type Catalog struct {
	Categories []Category `json:"categories"`
	Products   []Product  `json:"products"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Product struct {
	ID              int    `json:"id"`
	CategoryID      int    `json:"category"`
	Name            string `json:"name"`
	PriceParabool   int    `json:"priceParabool"`
	PriceGladiators int    `json:"priceGladiators"`
}
