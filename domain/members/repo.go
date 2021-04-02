package members

import (
	"encoding/json"
	"errors"

	"github.com/PotatoesFall/streepjes/shared"
	"go.etcd.io/bbolt"
)

var db *bbolt.DB

var ErrMemberNotFound = errors.New("member not found")

func getAll() []Member {
	members := []Member{}

	err := db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(shared.MembersBucket)
		return b.ForEach(func(_, v []byte) error {
			members = append(members, _unmarshal(v))
			return nil
		})
	})
	if err != nil {
		panic(err)
	}

	return members
}

// func _get(tx *bbolt.Tx, id int) (Member, error) {
// 	b := tx.Bucket(membersBucket)

// 	data := b.Get(shared.Itob(id))
// 	if data == nil {
// 		return Member{}, ErrMemberNotFound
// 	}

// 	return _unmarshal(data), nil
// }

// func _put(tx *bbolt.Tx, member Member) error {
// 	b := tx.Bucket(membersBucket)
// 	return b.Put(shared.Itob(member.ID), _marshal(member))
// }

func _unmarshal(data []byte) Member {
	var m Member
	err := json.Unmarshal(data, &m)
	if err != nil {
		panic(err)
	}
	return m
}

// func _marshal(member Member) []byte {
// 	data, err := json.Marshal(member)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return data
// }
