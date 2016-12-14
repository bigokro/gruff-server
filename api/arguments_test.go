package api

import (
	"encoding/json"
	"fmt"
	"github.com/bigokro/gruff-server/gruff"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestListArguments(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createArgument()
	u2 := createArgument()
	TESTDB.Create(&u1)
	TESTDB.Create(&u2)

	expectedResults, _ := json.Marshal([]gruff.Argument{u1, u2})

	r.GET("/api/arguments")
	res, _ := r.Run(Router())
	assert.Equal(t, string(expectedResults), res.Body.String())
	assert.Equal(t, http.StatusOK, res.Status())
}

func TestListArgumentsPagination(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createArgument()
	u2 := createArgument()
	TESTDB.Create(&u1)
	TESTDB.Create(&u2)

	r.GET("/api/arguments?start=0&limit=25")
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusOK, res.Status())
}

func TestGetArguments(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createArgument()
	TESTDB.Create(&u1)

	expectedResults, _ := json.Marshal(u1)

	r.GET(fmt.Sprintf("/api/arguments/%d", u1.ID))
	res, _ := r.Run(Router())
	assert.Equal(t, string(expectedResults), res.Body.String())
	assert.Equal(t, http.StatusOK, res.Status())
}

func TestCreateArguments(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createArgument()

	r.POST("/api/arguments")
	r.SetBody(u1)
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusCreated, res.Status())
}

func TestUpdateArguments(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createArgument()
	TESTDB.Create(&u1)

	r.PUT(fmt.Sprintf("/api/arguments/%d", u1.ID))
	r.SetBody(u1)
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusAccepted, res.Status())
}

func TestDeleteArguments(t *testing.T) {
	setup()
	defer teardown()
	r := New(Token)

	u1 := createArgument()
	TESTDB.Create(&u1)

	r.DELETE(fmt.Sprintf("/api/arguments/%d", u1.ID))
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusOK, res.Status())
}

func createArgument() gruff.Argument {
	a := gruff.Argument{
		Title:       "arguments",
		Description: "arguments",
		Type:        gruff.ARGUMENT_TYPE_PRO_TRUTH,
	}

	return a
}
