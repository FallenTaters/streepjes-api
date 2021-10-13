package buckets

import (
	"time"

	"github.com/FallenTaters/bbucket"
	"go.etcd.io/bbolt"
)

const path = "streepjes.db"

var DB *bbolt.DB

func Init() func() {
	database, err := bbolt.Open(path, 0666, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(err)
	}
	DB = database

	Orders = bbucket.New(DB, ordersBucketName)
	Users = bbucket.New(DB, usersBucketName)
	Members = bbucket.New(DB, membersBucketName)
	Categories = bbucket.New(DB, categoriesBucketName)
	Products = bbucket.New(DB, productsBucketName)

	return func() { _ = DB.Close() }
}

var (
	Orders     bbucket.Bucket
	Users      bbucket.Bucket
	Members    bbucket.Bucket
	Categories bbucket.Bucket
	Products   bbucket.Bucket

	ordersBucketName     = []byte("orders")
	usersBucketName      = []byte("users")
	membersBucketName    = []byte("members")
	categoriesBucketName = []byte("categories")
	productsBucketName   = []byte("products")
)
