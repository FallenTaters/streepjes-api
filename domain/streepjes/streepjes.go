package streepjes

import (
	"github.com/FallenTaters/streepjes-api/domain/members"
	"github.com/FallenTaters/streepjes-api/domain/orders"
	"github.com/FallenTaters/streepjes-api/domain/users"
)

func DeleteMember(id int) error {
	if orders.MemberHasUnpaidOrders(id) {
		return members.ErrUnpaidOrders
	}

	return members.ForceDelete(id)
}

func DeleteUser(username string) error {
	if orders.UserHasOpenOrders(username) {
		return users.ErrUserHasOpenOrders
	}

	return users.Delete(username)
}
