package orders

func AddOrder(order Order) error {
	return create(order)
}

func Filter(filter OrderFilter) ([]Order, error) {
	return filtered(func(o Order) bool {
		if filter.Club.Valid && filter.Club.Int64 != int64(o.Club) {
			return false
		}
		if filter.Bartender.Valid && filter.Bartender.String != o.Bartender {
			return false
		}
		return true
	})
}
