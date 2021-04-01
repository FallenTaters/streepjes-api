package orders

import (
	"database/sql"
)

func Init(database *sql.DB) {
	db = database
}

func AddOrder(order Order) error {
	return insertOrder(order)
}

func Get(filter Filter) []Order {
	return getOrders(filter)
}
