package model_test

import (
	"testing"

	"github.com/DaveSaah/risky/model"
	"github.com/stretchr/testify/assert"
)

func TestDBAlive(t *testing.T) {
	conn, err := model.InitDB()
	if err != nil {
		t.Fatal(err)
	}

	defer conn.Close()

	assert.NoError(t, conn.Ping())
}
