package catalog

import (
	"database/sql"
)

var db *sql.DB

const getCatalogQuery = `
	SELECT C.id, C.name, P.id, P.name, P.price_parabool, P.price_gladiators
	FROM category C
		INNER JOIN product P
			ON C.id = P.category_id
	;
`

func getCatalog() Catalog {
	rows, err := db.Query(getCatalogQuery)
	if err != nil {
		panic(err)
	}

	categories, products := []Category{}, []Product{}
	categorySeen := map[int]bool{}

	for rows.Next() {
		var category Category
		var product Product

		err = rows.Scan(
			&category.ID, &category.Name,
			&product.ID, &product.Name,
			&product.PriceParabool, &product.PriceGladiators,
		)
		if err != nil {
			panic(err)
		}
		product.CategoryID = category.ID

		if !categorySeen[category.ID] {
			categories = append(categories, category)
			categorySeen[category.ID] = true
		}
		products = append(products, product)
	}

	return Catalog{categories, products}
}
