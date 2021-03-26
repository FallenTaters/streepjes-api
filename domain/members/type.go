package members

import "github.com/PotatoesFall/streepjes/shared"

type Member struct {
	ID   int         `json:"id"`
	Club shared.Club `json:"club"`
	Name string      `json:"name"`
	Debt int         `json:"debt"`
}
