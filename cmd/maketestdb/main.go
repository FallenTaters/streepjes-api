package main

import (
	"encoding/json"
	"time"

	"github.com/PotatoesFall/streepjes/domain/catalog"
	"github.com/PotatoesFall/streepjes/domain/members"
	"github.com/PotatoesFall/streepjes/domain/orders"
	"github.com/PotatoesFall/streepjes/domain/users"
	"github.com/PotatoesFall/streepjes/shared"
	"github.com/PotatoesFall/streepjes/shared/migrate"
	"go.etcd.io/bbolt"
)

var db *bbolt.DB

const path = "streepjes.db"

func main() {
	getDB()
	initStuff()

	insertUsers()
	insertData()
}

func getDB() {
	database, err := bbolt.Open(path, 0666, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(err)
	}
	db = database

	err = migrate.Migrate(database)
	if err != nil {
		panic(err)
	}
}

var buckets = []string{
	"users",
	"categories",
	"products",
	"members",
	"orders",
}

func initStuff() {
	_ = db.Update(func(tx *bbolt.Tx) error {
		for _, bucket := range buckets {
			_ = tx.DeleteBucket([]byte(bucket))
		}
		return nil
	})

	catalog.Init(db)
	users.Init(db)
	members.Init(db)
	orders.Init(db)
	shared.CreateBuckets(db)
}

func insertUsers() {
	for _, u := range testUsers {
		err := users.Insert(u)
		if err != nil {
			panic(err)
		}
	}
}

func insertData() {
	err := db.Update(func(tx *bbolt.Tx) error {
		for _, bucket := range testData {
			b := tx.Bucket(bucket.Bucket)
			for _, pair := range bucket.Pairs {
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
