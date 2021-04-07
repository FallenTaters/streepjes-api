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

		if filter.Member.Valid && filter.Member.Int64 != int64(o.MemberID) {
			return false
		}

		if filter.Status.Valid && filter.Status.Int64 != int64(o.Status) {
			return false
		}

		return true
	})
}

func Delete(id int) error {
	return deleteByID(id)
}
