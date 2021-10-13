package main

import (
	"encoding/json"

	"github.com/FallenTaters/streepjes-api/domain/users"
	"github.com/FallenTaters/streepjes-api/repo/buckets"
	"go.etcd.io/bbolt"
)

func main() {
	close := buckets.Init()
	defer close()

	deleteData()
	insertUsers()
	insertData()
}

var bucketNames = []string{
	"users",
	"categories",
	"products",
	"members",
	"orders",
}

func deleteData() {
	_ = buckets.DB.Update(func(tx *bbolt.Tx) error {
		for _, bucket := range bucketNames {
			_ = tx.DeleteBucket([]byte(bucket))
			_, _ = tx.CreateBucket([]byte(bucket))
		}

		return nil
	})
}

func insertUsers() {
	for _, u := range testUsers {
		err := users.Put(u)
		if err != nil {
			panic(err)
		}
	}
}

func insertData() {
	err := buckets.DB.Update(func(tx *bbolt.Tx) error {
		for _, bucket := range testData {
			b := tx.Bucket(bucket.Bucket)
			for _, pair := range bucket.Pairs {
				b.NextSequence() //nolint: errcheck

				data, err := json.Marshal(pair.Value)
				if err != nil {
					panic(err)
				}

				err = b.Put(pair.Key, data)
				if err != nil {
					panic(err)
				}
			}
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
}

type bucketData struct {
	Bucket []byte
	Pairs  []keyValuePair
}

type keyValuePair struct {
	Key   []byte
	Value interface{}
}
