package main

import (
	"testing"

	"github.com/PotatoesFall/streepjes/cmd/streepjes/test"
	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	getDB()
	defer db.Close()

	_, err := db.Exec(string(test.MustGetFile(`testdata.sql`)))
	assert.NoError(t, err)
}
