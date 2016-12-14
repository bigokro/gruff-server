package api

import (
	"encoding/json"
	"fmt"
	"github.com/bigokro/gruff-server/gruff"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestListDebateOpinions(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createDebateOpinion(TESTDB)
	u2 := createDebateOpinion(TESTDB)
	TESTDB.Create(&u1)
	TESTDB.Create(&u2)

	expectedResults, _ := json.Marshal([]gruff.DebateOpinion{u1, u2})

	r.GET("/api/debate_opinions")
	res, _ := r.Run(Router())
	assert.Equal(t, string(expectedResults), res.Body.String())
	assert.Equal(t, http.StatusOK, res.Status())
}

func TestListDebateOpinionsPagination(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createDebateOpinion(TESTDB)
	u2 := createDebateOpinion(TESTDB)
	TESTDB.Create(&u1)
	TESTDB.Create(&u2)

	r.GET("/api/debate_opinions?start=0&limit=25")
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusOK, res.Status())
}

func TestGetDebateOpinions(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createDebateOpinion(TESTDB)
	TESTDB.Create(&u1)

	expectedResults, _ := json.Marshal(u1)

	r.GET(fmt.Sprintf("/api/debate_opinions/%d", u1.ID))
	res, _ := r.Run(Router())
	assert.Equal(t, string(expectedResults), res.Body.String())
	assert.Equal(t, http.StatusOK, res.Status())
}

func TestCreateDebateOpinions(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createDebateOpinion(TESTDB)

	r.POST("/api/debate_opinions")
	r.SetBody(u1)
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusCreated, res.Status())
}

func TestUpdateDebateOpinions(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createDebateOpinion(TESTDB)
	TESTDB.Create(&u1)

	r.PUT(fmt.Sprintf("/api/debate_opinions/%d", u1.ID))
	r.SetBody(u1)
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusAccepted, res.Status())
}

func TestDeleteDebateOpinions(t *testing.T) {
	setup()
	defer teardown()
	r := New(Token)

	u1 := createDebateOpinion(TESTDB)
	TESTDB.Create(&u1)

	r.DELETE(fmt.Sprintf("/api/debate_opinions/%d", u1.ID))
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusOK, res.Status())
}

func createDebateOpinion(db *gorm.DB) gruff.DebateOpinion {

	u := createUserAO(db)

	d := gruff.Debate{
		Title:       "Debate",
		Description: "Debate",
	}

	db.Create(&d)

	do := gruff.DebateOpinion{
		UserID:   uint64(u.ID),
		DebateID: uint64(d.ID),
		Truth:    10,
	}

	return do
}
