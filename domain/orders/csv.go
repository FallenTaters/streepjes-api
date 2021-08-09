package orders

import (
	"bytes"
	"encoding/csv"
	"strconv"

	"github.com/PotatoesFall/streepjes/domain/members"
	"github.com/PotatoesFall/streepjes/domain/users"
)

func MakeCSV(orders []Order) ([]byte, error) {
	csvFile := new(bytes.Buffer)
	w := csv.NewWriter(csvFile)

	err := w.Write([]string{"order id", "club", "bartender", "bartender name", "member id", "member name", "price", "date", "time", "status", "last updated", "contents"})
	if err != nil {
		panic(err)
	}

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
			member.Name = "unknown"
		} else if err != nil {
			panic(err)
		}

		orderDate := order.OrderTime.Format(`2006-01-02`)
		orderTime := order.OrderTime.Format(`15:04`)
		statusTime := order.StatusTime.Format(`2006-01-02`)

		err = w.Write([]string{
			omitempty(order.ID),
			order.Club.String(),
			order.Bartender,
			bartender.Name,
			omitempty(order.MemberID),
			member.Name,
			strconv.Itoa(order.Price),
			orderDate,
			orderTime,
			order.Status.String(),
			statusTime,
			order.Contents,
		})
		if err != nil {
			panic(err)
		}
	}

	w.Flush()
	return csvFile.Bytes(), nil
}

func omitempty(i int) string {
	if i == 0 {
		return ``
	}

	return strconv.Itoa(i)
}
