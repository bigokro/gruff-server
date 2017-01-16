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

func TestListClaimOpinions(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createClaimOpinion(TESTDB)
	u2 := createClaimOpinion(TESTDB)
	TESTDB.Create(&u1)
	TESTDB.Create(&u2)

	expectedResults, _ := json.Marshal([]gruff.ClaimOpinion{u1, u2})

	r.GET("/api/claim_opinions")
	res, _ := r.Run(Router())
	assert.Equal(t, string(expectedResults), res.Body.String())
	assert.Equal(t, http.StatusOK, res.Status())
}

func TestListClaimOpinionsPagination(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createClaimOpinion(TESTDB)
	u2 := createClaimOpinion(TESTDB)
	TESTDB.Create(&u1)
	TESTDB.Create(&u2)

	r.GET("/api/claim_opinions?start=0&limit=25")
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusOK, res.Status())
}

func TestGetClaimOpinions(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createClaimOpinion(TESTDB)
	TESTDB.Create(&u1)

	expectedResults, _ := json.Marshal(u1)

	r.GET(fmt.Sprintf("/api/claim_opinions/%d", u1.ID))
	res, _ := r.Run(Router())
	assert.Equal(t, string(expectedResults), res.Body.String())
	assert.Equal(t, http.StatusOK, res.Status())
}

func TestCreateClaimOpinions(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createClaimOpinion(TESTDB)

	r.POST("/api/claim_opinions")
	r.SetBody(u1)
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusCreated, res.Status())
}

func TestUpdateClaimOpinions(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createClaimOpinion(TESTDB)
	TESTDB.Create(&u1)

	r.PUT(fmt.Sprintf("/api/claim_opinions/%d", u1.ID))
	r.SetBody(u1)
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusAccepted, res.Status())
}

func TestDeleteClaimOpinions(t *testing.T) {
	setup()
	defer teardown()
	r := New(Token)

	u1 := createClaimOpinion(TESTDB)
	TESTDB.Create(&u1)

	r.DELETE(fmt.Sprintf("/api/claim_opinions/%d", u1.ID))
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusOK, res.Status())
}

func createClaimOpinion(db *gorm.DB) gruff.ClaimOpinion {

	u := createUserAO(db)

	d := gruff.Claim{
		Title:       "Claim",
		Description: "Claim",
	}

	db.Create(&d)

	do := gruff.ClaimOpinion{
		UserID:  uint64(u.ID),
		ClaimID: d.ID,
		Truth:   10,
	}

	return do
}
