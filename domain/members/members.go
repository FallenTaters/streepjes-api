package members

import (
	"go.etcd.io/bbolt"
)

func Init(database *bbolt.DB) {
	db = database
}

func GetAll() []Member {
	return getAll()
}
