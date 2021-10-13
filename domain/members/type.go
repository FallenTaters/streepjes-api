package members

import (
	"github.com/FallenTaters/streepjes-api/shared"
)

type Member struct {
	ID   int         `json:"id"`
	Club shared.Club `json:"club"`
	Name string      `json:"name"`
	Debt int         `json:"debt"`
}

func (m Member) Key() []byte {
	return shared.Itob(m.ID)
}
