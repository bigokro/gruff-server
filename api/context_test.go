package api

import (
	"encoding/json"
	"fmt"
	"github.com/bigokro/gruff-server/gruff"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestListContexts(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createContext()
	u2 := createContext()
	TESTDB.Create(&u1)
	TESTDB.Create(&u2)

	expectedResults, _ := json.Marshal([]gruff.Context{u1, u2})

	r.GET("/api/contexts")
	res, _ := r.Run(Router())
	assert.Equal(t, string(expectedResults), res.Body.String())
	assert.Equal(t, http.StatusOK, res.Status())
}

func TestListContextsPagination(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createContext()
	u2 := createContext()
	TESTDB.Create(&u1)
	TESTDB.Create(&u2)

	r.GET("/api/contexts?start=0&limit=25")
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusOK, res.Status())
}

func TestGetContexts(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createContext()
	TESTDB.Create(&u1)

	expectedResults, _ := json.Marshal(u1)

	r.GET(fmt.Sprintf("/api/contexts/%d", u1.ID))
	res, _ := r.Run(Router())
	assert.Equal(t, string(expectedResults), res.Body.String())
	assert.Equal(t, http.StatusOK, res.Status())
}

func TestCreateContexts(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createContext()

	r.POST("/api/contexts")
	r.SetBody(u1)
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusCreated, res.Status())
}

func TestUpdateContexts(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createContext()
	TESTDB.Create(&u1)

	r.PUT(fmt.Sprintf("/api/contexts/%d", u1.ID))
	r.SetBody(u1)
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusAccepted, res.Status())
}

func TestDeleteContexts(t *testing.T) {
	setup()
	defer teardown()
	r := New(Token)

	u1 := createContext()
	TESTDB.Create(&u1)

	r.DELETE(fmt.Sprintf("/api/contexts/%d", u1.ID))
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusOK, res.Status())
}

func createContext() gruff.Context {
	c := gruff.Context{
		Title:       "contexts",
		Description: "contexts",
	}

	return c
}
