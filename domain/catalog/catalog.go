package catalog

import (
	"time"
)

const cacheTime = time.Minute

func Get() (Catalog, error) {
	if time.Since(lastCatalog.Time) > cacheTime {
		c, err := getCatalog()
		if err != nil {
			return Catalog{}, err
		}
		lastCatalog.Catalog = c
		lastCatalog.Time = time.Now()
	}

	return lastCatalog.Catalog, nil
}

var lastCatalog struct {
	Catalog
	Time time.Time
}
