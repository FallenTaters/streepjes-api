package streepjes

import (
	"github.com/PotatoesFall/streepjes/domain/members"
	"github.com/PotatoesFall/streepjes/domain/orders"
)

func DeleteMember(id int) error {
	if orders.MemberHasUnpaidOrders(id) {
		return members.ErrUnpaidOrders
	}

	return members.ForceDelete(id)
}
