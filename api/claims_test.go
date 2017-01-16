package api

import (
	"encoding/json"
	"fmt"
	"github.com/bigokro/gruff-server/gruff"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestListClaims(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createClaim()
	u2 := createClaim()
	TESTDB.Create(&u1)
	TESTDB.Create(&u2)

	expectedResults, _ := json.Marshal([]gruff.Claim{u1, u2})

	r.GET("/api/claims")
	res, _ := r.Run(Router())
	assert.Equal(t, string(expectedResults), res.Body.String())
	assert.Equal(t, http.StatusOK, res.Status())
}

func TestListClaimsPagination(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createClaim()
	u2 := createClaim()
	TESTDB.Create(&u1)
	TESTDB.Create(&u2)

	r.GET("/api/claims?start=0&limit=25")
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusOK, res.Status())
}

func TestGetClaim(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	d1 := gruff.Claim{
		Title:       "Claim",
		Description: "This is a test Claim",
		Truth:       0.001,
	}
	d2 := gruff.Claim{
		Title:       "Another Claim",
		Description: "This is not the Claim you are looking for",
		Truth:       1.000,
	}
	d3 := gruff.Claim{Title: "Claim 3", Description: "Claim 3", Truth: 0.23094}
	d4 := gruff.Claim{Title: "Claim 4", Description: "Claim 4", Truth: 0.23094}
	d5 := gruff.Claim{Title: "Claim 5", Description: "Claim 5", Truth: 0.23094}
	d6 := gruff.Claim{Title: "Claim 6", Description: "Claim 6", Truth: 0.23094}
	d7 := gruff.Claim{Title: "Claim 7", Description: "Claim 7", Truth: 0.23094}
	d8 := gruff.Claim{Title: "Claim 8", Description: "Claim 8", Truth: 0.23094}
	d9 := gruff.Claim{Title: "Claim 9", Description: "Claim 9", Truth: 0.23094}
	TESTDB.Create(&d1)
	TESTDB.Create(&d2)
	TESTDB.Create(&d3)
	TESTDB.Create(&d4)
	TESTDB.Create(&d5)
	TESTDB.Create(&d6)
	TESTDB.Create(&d7)
	TESTDB.Create(&d8)
	TESTDB.Create(&d9)

	l1 := gruff.Link{Title: "A Link", Description: "What'd you expect?", Url: "http://site.com/page", ClaimID: d1.ID}
	l2 := gruff.Link{Title: "Another Link", Description: "What'd you expect?", Url: "http://site2.com/page", ClaimID: d1.ID}
	l3 := gruff.Link{Title: "An Irrelevant Link", Description: "What'd you expect?", Url: "http://site3.com/page", ClaimID: d2.ID}
	l4 := gruff.Link{Title: "Link 4", Description: "Link 4", Url: "http://site4.com/page", ClaimID: d4.ID}
	l5 := gruff.Link{Title: "Link 5", Description: "Link 5", Url: "http://site5.com/page", ClaimID: d5.ID}
	l6 := gruff.Link{Title: "Link 6", Description: "Link 6", Url: "http://site6.com/page", ClaimID: d6.ID}
	l7 := gruff.Link{Title: "Link 7", Description: "Link 7", Url: "http://site7.com/page", ClaimID: d7.ID}
	l8 := gruff.Link{Title: "Link 8", Description: "Link 8", Url: "http://site8.com/page", ClaimID: d8.ID}
	l9 := gruff.Link{Title: "Link 9", Description: "Link 9", Url: "http://site9.com/page", ClaimID: d9.ID}
	TESTDB.Create(&l1)
	TESTDB.Create(&l2)
	TESTDB.Create(&l3)
	TESTDB.Create(&l4)
	TESTDB.Create(&l5)
	TESTDB.Create(&l6)
	TESTDB.Create(&l7)
	TESTDB.Create(&l8)
	TESTDB.Create(&l9)

	c1 := gruff.Context{Title: "Test", Description: "The claim in question is related to a test"}
	c2 := gruff.Context{Title: "Valid", Description: "The claim in question is the one we want"}
	c3 := gruff.Context{Title: "Invalid", Description: "We don't want this"}
	TESTDB.Create(&c1)
	TESTDB.Create(&c2)
	TESTDB.Create(&c3)

	TESTDB.Exec("INSERT INTO claim_contexts (context_id, claim_id) VALUES (?, ?)", c1.ID, d1.ID)
	TESTDB.Exec("INSERT INTO claim_contexts (context_id, claim_id) VALUES (?, ?)", c2.ID, d1.ID)
	TESTDB.Exec("INSERT INTO claim_contexts (context_id, claim_id) VALUES (?, ?)", c1.ID, d2.ID)
	TESTDB.Exec("INSERT INTO claim_contexts (context_id, claim_id) VALUES (?, ?)", c3.ID, d2.ID)
	TESTDB.Exec("INSERT INTO claim_contexts (context_id, claim_id) VALUES (?, ?)", c3.ID, d3.ID)
	TESTDB.Exec("INSERT INTO claim_contexts (context_id, claim_id) VALUES (?, ?)", c3.ID, d4.ID)
	TESTDB.Exec("INSERT INTO claim_contexts (context_id, claim_id) VALUES (?, ?)", c3.ID, d5.ID)
	TESTDB.Exec("INSERT INTO claim_contexts (context_id, claim_id) VALUES (?, ?)", c3.ID, d6.ID)
	TESTDB.Exec("INSERT INTO claim_contexts (context_id, claim_id) VALUES (?, ?)", c3.ID, d7.ID)
	TESTDB.Exec("INSERT INTO claim_contexts (context_id, claim_id) VALUES (?, ?)", c3.ID, d8.ID)
	TESTDB.Exec("INSERT INTO claim_contexts (context_id, claim_id) VALUES (?, ?)", c3.ID, d9.ID)

	v1 := gruff.Value{Title: "Test", Description: "Testing is good"}
	v2 := gruff.Value{Title: "Correctness", Description: "We want correct code and tests"}
	v3 := gruff.Value{Title: "Completeness", Description: "The test suite should be complete enough to protect against all likely bugs"}
	TESTDB.Create(&v1)
	TESTDB.Create(&v2)
	TESTDB.Create(&v3)

	TESTDB.Exec("INSERT INTO claim_values (value_id, claim_id) VALUES (?, ?)", v1.ID, d1.ID)
	TESTDB.Exec("INSERT INTO claim_values (value_id, claim_id) VALUES (?, ?)", v2.ID, d1.ID)
	TESTDB.Exec("INSERT INTO claim_values (value_id, claim_id) VALUES (?, ?)", v1.ID, d2.ID)
	TESTDB.Exec("INSERT INTO claim_values (value_id, claim_id) VALUES (?, ?)", v3.ID, d2.ID)
	TESTDB.Exec("INSERT INTO claim_values (value_id, claim_id) VALUES (?, ?)", v3.ID, d3.ID)
	TESTDB.Exec("INSERT INTO claim_values (value_id, claim_id) VALUES (?, ?)", v3.ID, d4.ID)
	TESTDB.Exec("INSERT INTO claim_values (value_id, claim_id) VALUES (?, ?)", v3.ID, d5.ID)
	TESTDB.Exec("INSERT INTO claim_values (value_id, claim_id) VALUES (?, ?)", v3.ID, d6.ID)
	TESTDB.Exec("INSERT INTO claim_values (value_id, claim_id) VALUES (?, ?)", v3.ID, d7.ID)
	TESTDB.Exec("INSERT INTO claim_values (value_id, claim_id) VALUES (?, ?)", v3.ID, d8.ID)
	TESTDB.Exec("INSERT INTO claim_values (value_id, claim_id) VALUES (?, ?)", v3.ID, d9.ID)

	t1 := gruff.Tag{Title: "Test"}
	t2 := gruff.Tag{Title: "Valid"}
	t3 := gruff.Tag{Title: "Invalid"}
	TESTDB.Create(&t1)
	TESTDB.Create(&t2)
	TESTDB.Create(&t3)

	TESTDB.Exec("INSERT INTO claim_tags (tag_id, claim_id) VALUES (?, ?)", t1.ID, d1.ID)
	TESTDB.Exec("INSERT INTO claim_tags (tag_id, claim_id) VALUES (?, ?)", t2.ID, d1.ID)
	TESTDB.Exec("INSERT INTO claim_tags (tag_id, claim_id) VALUES (?, ?)", t1.ID, d2.ID)
	TESTDB.Exec("INSERT INTO claim_tags (tag_id, claim_id) VALUES (?, ?)", t3.ID, d2.ID)
	TESTDB.Exec("INSERT INTO claim_tags (tag_id, claim_id) VALUES (?, ?)", t3.ID, d3.ID)
	TESTDB.Exec("INSERT INTO claim_tags (tag_id, claim_id) VALUES (?, ?)", t3.ID, d4.ID)
	TESTDB.Exec("INSERT INTO claim_tags (tag_id, claim_id) VALUES (?, ?)", t3.ID, d5.ID)
	TESTDB.Exec("INSERT INTO claim_tags (tag_id, claim_id) VALUES (?, ?)", t3.ID, d6.ID)
	TESTDB.Exec("INSERT INTO claim_tags (tag_id, claim_id) VALUES (?, ?)", t3.ID, d7.ID)
	TESTDB.Exec("INSERT INTO claim_tags (tag_id, claim_id) VALUES (?, ?)", t3.ID, d8.ID)
	TESTDB.Exec("INSERT INTO claim_tags (tag_id, claim_id) VALUES (?, ?)", t3.ID, d9.ID)

	d1IDNull := gruff.NullableUUID{d1.ID}
	d2IDNull := gruff.NullableUUID{d2.ID}
	d3IDNull := gruff.NullableUUID{d3.ID}
	d4IDNull := gruff.NullableUUID{d4.ID}
	a3 := gruff.Argument{TargetClaimID: &d1IDNull, ClaimID: d3.ID, Type: gruff.ARGUMENT_TYPE_PRO_TRUTH, Title: "Argument 3", Relevance: 0.2309, Impact: 0.0293}
	a4 := gruff.Argument{TargetClaimID: &d1IDNull, ClaimID: d4.ID, Type: gruff.ARGUMENT_TYPE_CON_TRUTH, Title: "Argument 4", Relevance: 0.29, Impact: 0.9823}
	a5 := gruff.Argument{TargetClaimID: &d1IDNull, ClaimID: d5.ID, Type: gruff.ARGUMENT_TYPE_PRO_TRUTH, Title: "Argument 5", Relevance: 0.4893, Impact: 0.100}
	a6 := gruff.Argument{TargetClaimID: &d2IDNull, ClaimID: d6.ID, Type: gruff.ARGUMENT_TYPE_PRO_TRUTH, Title: "Argument 6", Relevance: 0.438, Impact: 0.2398}
	a7 := gruff.Argument{TargetClaimID: &d2IDNull, ClaimID: d7.ID, Type: gruff.ARGUMENT_TYPE_CON_TRUTH, Title: "Argument 7", Relevance: 0.2398, Impact: 0.120}
	a8 := gruff.Argument{TargetArgumentID: &d3IDNull, ClaimID: d8.ID, Type: gruff.ARGUMENT_TYPE_PRO_RELEVANCE, Title: "Argument 8", Relevance: 0.9, Impact: 0.9823}
	a9 := gruff.Argument{TargetArgumentID: &d3IDNull, ClaimID: d9.ID, Type: gruff.ARGUMENT_TYPE_CON_RELEVANCE, Title: "Argument 9", Relevance: 0.2398, Impact: 0.83}
	a10 := gruff.Argument{TargetClaimID: &d4IDNull, ClaimID: d3.ID, Type: gruff.ARGUMENT_TYPE_CON_IMPACT, Title: "Argument 10", Relevance: 0.2398, Impact: 0.83}
	TESTDB.Create(&a3)
	TESTDB.Create(&a4)
	TESTDB.Create(&a5)
	TESTDB.Create(&a6)
	TESTDB.Create(&a7)
	TESTDB.Create(&a8)
	TESTDB.Create(&a9)
	TESTDB.Create(&a10)

	a3.Claim = &d3
	a4.Claim = &d4
	a5.Claim = &d5
	a6.Claim = &d6
	a7.Claim = &d7
	a8.Claim = &d8
	a9.Claim = &d9
	a10.Claim = &d3

	db := TESTDB
	db = db.Preload("Links")
	db = db.Preload("Contexts")
	db = db.Preload("Values")
	db = db.Preload("Tags")
	db.Where("id = ?", d1.ID).First(&d1)

	d1.ProTruth = []gruff.Argument{a3, a5}
	d1.ConTruth = []gruff.Argument{a4}

	expectedResults, _ := json.Marshal(d1)

	r.GET(fmt.Sprintf("/api/claims/%s", d1.ID.String()))
	res, _ := r.Run(Router())
	assert.Equal(t, string(expectedResults), res.Body.String())
	assert.Equal(t, http.StatusOK, res.Status())
}

func TestCreateClaim(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createClaim()

	r.POST("/api/claims")
	r.SetBody(u1)
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusCreated, res.Status())
}

func TestUpdateClaim(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createClaim()
	TESTDB.Create(&u1)

	r.PUT(fmt.Sprintf("/api/claims/%s", u1.ID))
	r.SetBody(u1)
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusAccepted, res.Status())
}

func TestDeleteClaim(t *testing.T) {
	setup()
	defer teardown()
	r := New(Token)

	u1 := createClaim()
	TESTDB.Create(&u1)

	r.DELETE(fmt.Sprintf("/api/claims/%s", u1.ID))
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusOK, res.Status())
}

func createClaim() gruff.Claim {
	c := gruff.Claim{
		Title:       "Claim",
		Description: "Claim",
	}

	return c
}
