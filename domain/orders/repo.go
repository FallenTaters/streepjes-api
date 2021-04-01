package orders

import (
	"database/sql"

	"github.com/PotatoesFall/streepjes/shared/sql/wherebuilder"
)

var db *sql.DB

var selectBase = `SELECT id, club, bartender_id, member_id, contents, price, order_datetime, status, status_datetime FROM "order"`

var insertOrderQ = `INSERT INTO "order"(club, bartender_id, member_id, contents, price, order_datetime, status, status_datetime)
VALUES($1, $2, $3, $4, $5, $6, $7, $8);`

func insertOrder(order Order) error {
	_, err := db.Exec(insertOrderQ, order.Club, order.BartenderID, order.MemberID, order.Contents, order.Price, order.OrderTime, order.Status, order.StatusTime)
	return err
}

func mustScanOrder(rows *sql.Rows) Order {
	var o Order
	err := rows.Scan(&o.ID, &o.Club, &o.BartenderID, &o.MemberID, &o.Contents, &o.Price, &o.OrderTime, &o.Status, &o.StatusTime)
	if err != nil {
		panic(err)
	}
	return o
}

func getOrders(filter Filter) []Order {
	wb := wherebuilder.New(selectBase)

	if filter.BartenderID.Valid {
		wb.Where(`bartender_id = ` + filter.BartenderID.String())
	}
	if filter.Club.Valid {
		wb.Where(`club = ` + filter.Club.String())
	}

	rows, err := db.Query(wb.Query())
	if err != nil {
		panic(err)
	}
	defer rows.Close() //nolint: errcheck

	var orders []Order
	for rows.Next() {
		orders = append(orders, mustScanOrder(rows))
	}

	return orders
}
