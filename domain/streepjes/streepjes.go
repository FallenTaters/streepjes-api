package streepjes

import (
	"github.com/FallenTaters/streepjes-api/domain/members"
	"github.com/FallenTaters/streepjes-api/domain/orders"
)

func DeleteMember(id int) error {
	if orders.MemberHasUnpaidOrders(id) {
		return members.ErrUnpaidOrders
	}

	return members.ForceDelete(id)
}
