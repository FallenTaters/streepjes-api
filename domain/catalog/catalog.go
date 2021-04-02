package catalog

import (
	"time"

	"go.etcd.io/bbolt"
)

const cacheTime = time.Minute

func Init(database *bbolt.DB) {
	db = database
}

func Get() Catalog {
	if time.Since(lastCatalog.Time) > cacheTime {
		lastCatalog.Catalog = getCatalog()
		lastCatalog.Time = time.Now()
	}

	return lastCatalog.Catalog
}

var lastCatalog struct {
	Catalog
	Time time.Time
}
