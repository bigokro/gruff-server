package api

import (
	"encoding/json"
	"fmt"
	"github.com/bigokro/gruff-server/gruff"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestListDebates(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createDebate()
	u2 := createDebate()
	TESTDB.Create(&u1)
	TESTDB.Create(&u2)

	expectedResults, _ := json.Marshal([]gruff.Debate{u1, u2})

	r.GET("/api/debates")
	res, _ := r.Run(Router())
	assert.Equal(t, string(expectedResults), res.Body.String())
	assert.Equal(t, http.StatusOK, res.Status())
}

func TestListDebatesPagination(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createDebate()
	u2 := createDebate()
	TESTDB.Create(&u1)
	TESTDB.Create(&u2)

	r.GET("/api/debates?start=0&limit=25")
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusOK, res.Status())
}

func TestGetDebates(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createDebate()
	TESTDB.Create(&u1)

	expectedResults, _ := json.Marshal(u1)

	r.GET(fmt.Sprintf("/api/debates/%d", u1.ID))
	res, _ := r.Run(Router())
	assert.Equal(t, string(expectedResults), res.Body.String())
	assert.Equal(t, http.StatusOK, res.Status())
}

func TestCreateDebates(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createDebate()

	r.POST("/api/debates")
	r.SetBody(u1)
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusCreated, res.Status())
}

func TestUpdateDebates(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createDebate()
	TESTDB.Create(&u1)

	r.PUT(fmt.Sprintf("/api/debates/%d", u1.ID))
	r.SetBody(u1)
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusAccepted, res.Status())
}

func TestDeleteDebates(t *testing.T) {
	setup()
	defer teardown()
	r := New(Token)

	u1 := createDebate()
	TESTDB.Create(&u1)

	r.DELETE(fmt.Sprintf("/api/debates/%d", u1.ID))
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusOK, res.Status())
}

func createDebate() gruff.Debate {
	c := gruff.Debate{
		Title:       "Debate",
		Description: "Debate",
	}

	return c
}
