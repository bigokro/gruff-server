package gruff

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateDebate(t *testing.T) {
	setupDB()
	defer teardownDB()

	d := Debate{
		Title:       "The first debate!",
		Description: "A description",
		Truth:       87.55,
	}
	TESTDB.Create(&d)

	assert.True(t, d.ID != ZERO_UUID)
}
