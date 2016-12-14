package api

import (
	"encoding/json"
	"fmt"
	"github.com/bigokro/gruff-server/gruff"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestListLinks(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createLink()
	u2 := createLink()
	TESTDB.Create(&u1)
	TESTDB.Create(&u2)

	expectedResults, _ := json.Marshal([]gruff.Link{u1, u2})

	r.GET("/api/links")
	res, _ := r.Run(Router())
	assert.Equal(t, string(expectedResults), res.Body.String())
	assert.Equal(t, http.StatusOK, res.Status())
}

func TestListLinksPagination(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createLink()
	u2 := createLink()
	TESTDB.Create(&u1)
	TESTDB.Create(&u2)

	r.GET("/api/links?start=0&limit=25")
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusOK, res.Status())
}

func TestGetLinks(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createLink()
	TESTDB.Create(&u1)

	expectedResults, _ := json.Marshal(u1)

	r.GET(fmt.Sprintf("/api/links/%d", u1.ID))
	res, _ := r.Run(Router())
	assert.Equal(t, string(expectedResults), res.Body.String())
	assert.Equal(t, http.StatusOK, res.Status())
}

func TestCreateLinks(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createLink()

	r.POST("/api/links")
	r.SetBody(u1)
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusCreated, res.Status())
}

func TestUpdateLinks(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createLink()
	TESTDB.Create(&u1)

	r.PUT(fmt.Sprintf("/api/links/%d", u1.ID))
	r.SetBody(u1)
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusAccepted, res.Status())
}

func TestDeleteLinks(t *testing.T) {
	setup()
	defer teardown()
	r := New(Token)

	u1 := createLink()
	TESTDB.Create(&u1)

	r.DELETE(fmt.Sprintf("/api/links/%d", u1.ID))
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusOK, res.Status())
}

func createLink() gruff.Link {
	l := gruff.Link{
		Title:       "Links",
		Description: "Links",
		Url:         "www.gruff.org",
	}

	return l
}
