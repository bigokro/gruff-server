package api

import (
	"encoding/json"
	"fmt"
	"github.com/bigokro/gruff-server/gruff"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestListTags(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createTag()
	u2 := createTag()
	TESTDB.Create(&u1)
	TESTDB.Create(&u2)

	expectedResults, _ := json.Marshal([]gruff.Tag{u1, u2})

	r.GET("/api/tags")
	res, _ := r.Run(Router())
	assert.Equal(t, string(expectedResults), res.Body.String())
	assert.Equal(t, http.StatusOK, res.Status())
}

func TestListTagsPagination(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createTag()
	u2 := createTag()
	TESTDB.Create(&u1)
	TESTDB.Create(&u2)

	r.GET("/api/tags?start=0&limit=25")
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusOK, res.Status())
}

func TestGetTags(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createTag()
	TESTDB.Create(&u1)

	expectedResults, _ := json.Marshal(u1)

	r.GET(fmt.Sprintf("/api/tags/%d", u1.ID))
	res, _ := r.Run(Router())
	assert.Equal(t, string(expectedResults), res.Body.String())
	assert.Equal(t, http.StatusOK, res.Status())
}

func TestCreateTags(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createTag()

	r.POST("/api/tags")
	r.SetBody(u1)
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusCreated, res.Status())
}

func TestUpdateTags(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createTag()
	TESTDB.Create(&u1)

	r.PUT(fmt.Sprintf("/api/tags/%d", u1.ID))
	r.SetBody(u1)
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusAccepted, res.Status())
}

func TestDeleteTags(t *testing.T) {
	setup()
	defer teardown()
	r := New(Token)

	u1 := createTag()
	TESTDB.Create(&u1)

	r.DELETE(fmt.Sprintf("/api/tags/%d", u1.ID))
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusOK, res.Status())
}

func createTag() gruff.Tag {
	t := gruff.Tag{
		Title: "Tag",
	}

	return t
}
