package orders

import (
	"fmt"

	"github.com/PotatoesFall/streepjes/domain/members"
	"github.com/PotatoesFall/streepjes/domain/users"
)

func MakeCSV(orders []Order) ([]byte, error) {
	data := "order id, club, bartender, bartender name, member id, member name, contents, price, time, status, last updated\n"

	for _, order := range orders {
		bartender, err := users.Get(order.Bartender)
		if err == users.ErrUserNotFound {
			bartender.Name = "unknown"
		}
		if err != nil {
			panic(err)
		}

		member, err := members.Get(order.MemberID)
		if err == members.ErrMemberNotFound {
			bartender.Name = "unknown"
		}
		if err != nil {
			panic(err)
		}

		orderTime := order.OrderTime.Format(`2006-01-02`)
		statusTime := order.StatusTime.Format(`2006-01-02`)

		data += fmt.Sprintf("%d, %s, %s, %s, %d, %s %s, %d, %s, %s, %s\n", order.ID, order.Club, order.Bartender, bartender.Name, order.MemberID, member.Name, order.Contents, order.Price, orderTime, order.Status, statusTime)
	}

	return []byte(data), nil
}
