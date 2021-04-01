package orders

import (
	"database/sql"
	"time"
)

type OrderStatus int

const (
	OrderStatusOpen      OrderStatus = iota // Open
	OrderStatusFinished                     // Finished
	OrderStatusCancelled                    // Cancelled
)

//go:generate enumer -linecomment -type OrderStatus

type Order struct {
	ID          int           `json:"id"`
	BartenderID int           `json:"bartenderId"`
	MemberID    sql.NullInt32 `json:"memberId"`
	Contents    []byte        `json:"contents"`
	Price       int           `json:"price"`
	OrderTime   time.Time     `json:"orderDate"`
	Status      OrderStatus   `json:"status"`
	StatusTime  sql.NullTime  `json:"statusDate"`
}
