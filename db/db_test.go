package db_test

import (
	"testing"

	"github.com/DaveSaah/risky/db"
	"github.com/stretchr/testify/assert"
)

func TestDBAlive(t *testing.T) {
	conn, err := db.InitDB()
	if err != nil {
		t.Fatal(err)
	}

	defer conn.Close()

	assert.NoError(t, conn.Ping())
}
