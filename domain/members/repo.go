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

func get(id int) (Member, error) {
	var member Member
	return member, buckets.Members.Get(buckets.Itob(id), &member)
}

func updateMember(member Member) error {
	return buckets.Members.Update(buckets.Itob(member.ID), &Member{}, func(ptr interface{}) (object interface{}, err error) {
		return member, nil
	})
}

func addMember(member Member) error {
	return buckets.Members.Create(buckets.Itob(buckets.Members.NextSequence()), member)
}

func deleteMember(id int) error {
	return buckets.Members.Delete(buckets.Itob(id))
}
