package orders

import (
	"time"

	"github.com/PotatoesFall/streepjes/shared"
	"github.com/PotatoesFall/streepjes/shared/null"
)

type OrderStatus int

const (
	OrderStatusOpen      OrderStatus = iota + 1 // Open
	OrderStatusBilled                           // Billed
	OrderStatusPaid                             // Paid
	OrderStatusCancelled                        // Cancelled
)

//go:generate enumer -json -linecomment -type OrderStatus

type Order struct {
	ID          int         `json:"id"`
	Club        shared.Club `json:"club"`
	BartenderID int         `json:"bartenderId"`
	MemberID    int         `json:"memberId"`
	Contents    string      `json:"contents"`
	Price       int         `json:"price"`
	OrderTime   time.Time   `json:"orderDate"`
	Status      OrderStatus `json:"status"`
	StatusTime  time.Time   `json:"statusDate"`
}

type Filter struct {
	Club        null.Int `json:"club"`
	BartenderID null.Int `json:"bartenderId"`
}
