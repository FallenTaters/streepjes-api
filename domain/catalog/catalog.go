package catalog

import (
	"database/sql"
	"time"
)

const cacheTime = time.Minute

func Init(database *sql.DB) {
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
