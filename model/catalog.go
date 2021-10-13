package model

import "github.com/FallenTaters/streepjes-api/shared"

type Catalog struct {
	Categories []Category `json:"categories"`
	Products   []Product  `json:"products"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c Category) Key() []byte {
	return shared.Itob(c.ID)
}

type Product struct {
	ID              int    `json:"id"`
	CategoryID      int    `json:"category"`
	Name            string `json:"name"`
	PriceParabool   int    `json:"priceParabool"`
	PriceGladiators int    `json:"priceGladiators"`
}

func (p Product) Key() []byte {
	return shared.Itob(p.ID)
}
