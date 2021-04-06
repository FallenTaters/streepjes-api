package members

import (
	"errors"

	"github.com/PotatoesFall/streepjes/shared/buckets"
)

var ErrMemberNotFound = errors.New("member not found")

func getAll() ([]Member, error) {
	members := []Member{}

	return members, buckets.Members.GetAll(&Member{}, func(ptr interface{}) error {
		members = append(members, *ptr.(*Member))
		return nil
	})
}
