package orders

import (
	"time"

	"github.com/FallenTaters/streepjes-api/shared"
	"github.com/FallenTaters/streepjes-api/shared/buckets"
	"github.com/FallenTaters/streepjes-api/shared/null"
)

type OrderType int

const (
	OrderTypeBilled    OrderType = iota + 1 // Billed
	OrderTypePaid                           // Paid
	OrderTypeCancelled                      // Cancelled
)

//go:generate enumer -json -linecomment -type OrderType

type Order struct {
	ID        int         `json:"id"`
	Club      shared.Club `json:"club"`
	Bartender string      `json:"bartender"`
	MemberID  int         `json:"memberId"`
	Contents  string      `json:"contents"`
	Price     int         `json:"price"`
	OrderTime time.Time   `json:"orderDate"`
	Status    OrderType   `json:"status"`
	UpdatedAt time.Time   `json:"statusDate"`
}

func (o Order) Key() []byte {
	return buckets.Itob(o.ID)
}

func (o Order) IsEditable() bool {
	return time.Since(o.OrderTime).Hours()/24 < 7
}

type OrderFilter struct {
	Club      null.Int
	Bartender null.String
	Member    null.Int
	Status    []int
	Month     *time.Time
}
