package streepjes

import (
	"github.com/PotatoesFall/streepjes/domain/members"
	"github.com/PotatoesFall/streepjes/domain/orders"
	"github.com/PotatoesFall/streepjes/domain/users"
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
