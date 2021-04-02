package shared

import (
	"encoding/binary"

	"go.etcd.io/bbolt"
)

var (
	UsersBucket      = []byte("users")
	OrdersBucket     = []byte("orders")
	MembersBucket    = []byte("members")
	CategoriesBucket = []byte("categories")
	ProductsBucket   = []byte("products")
)

func CreateBuckets(db *bbolt.DB) {
	createBuckets(db, UsersBucket, OrdersBucket, MembersBucket, CategoriesBucket, ProductsBucket)
}

func createBuckets(db *bbolt.DB, buckets ...[]byte) {
	err := db.Update(func(tx *bbolt.Tx) error {
		for _, bucket := range buckets {
			_, err := tx.CreateBucketIfNotExists(bucket)
			if err != nil {
				panic(err)
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func Itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
