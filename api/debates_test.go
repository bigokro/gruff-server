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

func TestGetDebate(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	d1 := gruff.Debate{
		Title:       "Debate",
		Description: "This is a test Debate",
		Truth:       0.001,
	}
	d2 := gruff.Debate{
		Title:       "Another Debate",
		Description: "This is not the Debate you are looking for",
		Truth:       1.000,
	}
	d3 := gruff.Debate{Title: "Debate 3", Description: "Debate 3", Truth: 0.23094}
	d4 := gruff.Debate{Title: "Debate 4", Description: "Debate 4", Truth: 0.23094}
	d5 := gruff.Debate{Title: "Debate 5", Description: "Debate 5", Truth: 0.23094}
	d6 := gruff.Debate{Title: "Debate 6", Description: "Debate 6", Truth: 0.23094}
	d7 := gruff.Debate{Title: "Debate 7", Description: "Debate 7", Truth: 0.23094}
	d8 := gruff.Debate{Title: "Debate 8", Description: "Debate 8", Truth: 0.23094}
	d9 := gruff.Debate{Title: "Debate 9", Description: "Debate 9", Truth: 0.23094}
	TESTDB.Create(&d1)
	TESTDB.Create(&d2)
	TESTDB.Create(&d3)
	TESTDB.Create(&d4)
	TESTDB.Create(&d5)
	TESTDB.Create(&d6)
	TESTDB.Create(&d7)
	TESTDB.Create(&d8)
	TESTDB.Create(&d9)

	l1 := gruff.Link{Title: "A Link", Description: "What'd you expect?", Url: "http://site.com/page", DebateID: d1.ID}
	l2 := gruff.Link{Title: "Another Link", Description: "What'd you expect?", Url: "http://site2.com/page", DebateID: d1.ID}
	l3 := gruff.Link{Title: "An Irrelevant Link", Description: "What'd you expect?", Url: "http://site3.com/page", DebateID: d2.ID}
	l4 := gruff.Link{Title: "Link 4", Description: "Link 4", Url: "http://site4.com/page", DebateID: d4.ID}
	l5 := gruff.Link{Title: "Link 5", Description: "Link 5", Url: "http://site5.com/page", DebateID: d5.ID}
	l6 := gruff.Link{Title: "Link 6", Description: "Link 6", Url: "http://site6.com/page", DebateID: d6.ID}
	l7 := gruff.Link{Title: "Link 7", Description: "Link 7", Url: "http://site7.com/page", DebateID: d7.ID}
	l8 := gruff.Link{Title: "Link 8", Description: "Link 8", Url: "http://site8.com/page", DebateID: d8.ID}
	l9 := gruff.Link{Title: "Link 9", Description: "Link 9", Url: "http://site9.com/page", DebateID: d9.ID}
	TESTDB.Create(&l1)
	TESTDB.Create(&l2)
	TESTDB.Create(&l3)
	TESTDB.Create(&l4)
	TESTDB.Create(&l5)
	TESTDB.Create(&l6)
	TESTDB.Create(&l7)
	TESTDB.Create(&l8)
	TESTDB.Create(&l9)

	c1 := gruff.Context{Title: "Test", Description: "The debate in question is related to a test"}
	c2 := gruff.Context{Title: "Valid", Description: "The debate in question is the one we want"}
	c3 := gruff.Context{Title: "Invalid", Description: "We don't want this"}
	TESTDB.Create(&c1)
	TESTDB.Create(&c2)
	TESTDB.Create(&c3)

	TESTDB.Exec("INSERT INTO debate_contexts (context_id, debate_id) VALUES (?, ?)", c1.ID, d1.ID)
	TESTDB.Exec("INSERT INTO debate_contexts (context_id, debate_id) VALUES (?, ?)", c2.ID, d1.ID)
	TESTDB.Exec("INSERT INTO debate_contexts (context_id, debate_id) VALUES (?, ?)", c1.ID, d2.ID)
	TESTDB.Exec("INSERT INTO debate_contexts (context_id, debate_id) VALUES (?, ?)", c3.ID, d2.ID)
	TESTDB.Exec("INSERT INTO debate_contexts (context_id, debate_id) VALUES (?, ?)", c3.ID, d3.ID)
	TESTDB.Exec("INSERT INTO debate_contexts (context_id, debate_id) VALUES (?, ?)", c3.ID, d4.ID)
	TESTDB.Exec("INSERT INTO debate_contexts (context_id, debate_id) VALUES (?, ?)", c3.ID, d5.ID)
	TESTDB.Exec("INSERT INTO debate_contexts (context_id, debate_id) VALUES (?, ?)", c3.ID, d6.ID)
	TESTDB.Exec("INSERT INTO debate_contexts (context_id, debate_id) VALUES (?, ?)", c3.ID, d7.ID)
	TESTDB.Exec("INSERT INTO debate_contexts (context_id, debate_id) VALUES (?, ?)", c3.ID, d8.ID)
	TESTDB.Exec("INSERT INTO debate_contexts (context_id, debate_id) VALUES (?, ?)", c3.ID, d9.ID)

	v1 := gruff.Value{Title: "Test", Description: "Testing is good"}
	v2 := gruff.Value{Title: "Correctness", Description: "We want correct code and tests"}
	v3 := gruff.Value{Title: "Completeness", Description: "The test suite should be complete enough to protect against all likely bugs"}
	TESTDB.Create(&v1)
	TESTDB.Create(&v2)
	TESTDB.Create(&v3)

	TESTDB.Exec("INSERT INTO debate_values (value_id, debate_id) VALUES (?, ?)", v1.ID, d1.ID)
	TESTDB.Exec("INSERT INTO debate_values (value_id, debate_id) VALUES (?, ?)", v2.ID, d1.ID)
	TESTDB.Exec("INSERT INTO debate_values (value_id, debate_id) VALUES (?, ?)", v1.ID, d2.ID)
	TESTDB.Exec("INSERT INTO debate_values (value_id, debate_id) VALUES (?, ?)", v3.ID, d2.ID)
	TESTDB.Exec("INSERT INTO debate_values (value_id, debate_id) VALUES (?, ?)", v3.ID, d3.ID)
	TESTDB.Exec("INSERT INTO debate_values (value_id, debate_id) VALUES (?, ?)", v3.ID, d4.ID)
	TESTDB.Exec("INSERT INTO debate_values (value_id, debate_id) VALUES (?, ?)", v3.ID, d5.ID)
	TESTDB.Exec("INSERT INTO debate_values (value_id, debate_id) VALUES (?, ?)", v3.ID, d6.ID)
	TESTDB.Exec("INSERT INTO debate_values (value_id, debate_id) VALUES (?, ?)", v3.ID, d7.ID)
	TESTDB.Exec("INSERT INTO debate_values (value_id, debate_id) VALUES (?, ?)", v3.ID, d8.ID)
	TESTDB.Exec("INSERT INTO debate_values (value_id, debate_id) VALUES (?, ?)", v3.ID, d9.ID)

	t1 := gruff.Tag{Title: "Test"}
	t2 := gruff.Tag{Title: "Valid"}
	t3 := gruff.Tag{Title: "Invalid"}
	TESTDB.Create(&t1)
	TESTDB.Create(&t2)
	TESTDB.Create(&t3)

	TESTDB.Exec("INSERT INTO debate_tags (tag_id, debate_id) VALUES (?, ?)", t1.ID, d1.ID)
	TESTDB.Exec("INSERT INTO debate_tags (tag_id, debate_id) VALUES (?, ?)", t2.ID, d1.ID)
	TESTDB.Exec("INSERT INTO debate_tags (tag_id, debate_id) VALUES (?, ?)", t1.ID, d2.ID)
	TESTDB.Exec("INSERT INTO debate_tags (tag_id, debate_id) VALUES (?, ?)", t3.ID, d2.ID)
	TESTDB.Exec("INSERT INTO debate_tags (tag_id, debate_id) VALUES (?, ?)", t3.ID, d3.ID)
	TESTDB.Exec("INSERT INTO debate_tags (tag_id, debate_id) VALUES (?, ?)", t3.ID, d4.ID)
	TESTDB.Exec("INSERT INTO debate_tags (tag_id, debate_id) VALUES (?, ?)", t3.ID, d5.ID)
	TESTDB.Exec("INSERT INTO debate_tags (tag_id, debate_id) VALUES (?, ?)", t3.ID, d6.ID)
	TESTDB.Exec("INSERT INTO debate_tags (tag_id, debate_id) VALUES (?, ?)", t3.ID, d7.ID)
	TESTDB.Exec("INSERT INTO debate_tags (tag_id, debate_id) VALUES (?, ?)", t3.ID, d8.ID)
	TESTDB.Exec("INSERT INTO debate_tags (tag_id, debate_id) VALUES (?, ?)", t3.ID, d9.ID)

	d1IDNull := gruff.NullableUUID{d1.ID}
	d2IDNull := gruff.NullableUUID{d2.ID}
	d3IDNull := gruff.NullableUUID{d3.ID}
	d4IDNull := gruff.NullableUUID{d4.ID}
	a3 := gruff.Argument{TargetDebateID: &d1IDNull, DebateID: d3.ID, Type: gruff.ARGUMENT_TYPE_PRO_TRUTH, Title: "Argument 3", Relevance: 0.2309, Impact: 0.0293}
	a4 := gruff.Argument{TargetDebateID: &d1IDNull, DebateID: d4.ID, Type: gruff.ARGUMENT_TYPE_CON_TRUTH, Title: "Argument 4", Relevance: 0.29, Impact: 0.9823}
	a5 := gruff.Argument{TargetDebateID: &d1IDNull, DebateID: d5.ID, Type: gruff.ARGUMENT_TYPE_PRO_TRUTH, Title: "Argument 5", Relevance: 0.4893, Impact: 0.100}
	a6 := gruff.Argument{TargetDebateID: &d2IDNull, DebateID: d6.ID, Type: gruff.ARGUMENT_TYPE_PRO_TRUTH, Title: "Argument 6", Relevance: 0.438, Impact: 0.2398}
	a7 := gruff.Argument{TargetDebateID: &d2IDNull, DebateID: d7.ID, Type: gruff.ARGUMENT_TYPE_CON_TRUTH, Title: "Argument 7", Relevance: 0.2398, Impact: 0.120}
	a8 := gruff.Argument{TargetArgumentID: &d3IDNull, DebateID: d8.ID, Type: gruff.ARGUMENT_TYPE_PRO_RELEVANCE, Title: "Argument 8", Relevance: 0.9, Impact: 0.9823}
	a9 := gruff.Argument{TargetArgumentID: &d3IDNull, DebateID: d9.ID, Type: gruff.ARGUMENT_TYPE_CON_RELEVANCE, Title: "Argument 9", Relevance: 0.2398, Impact: 0.83}
	a10 := gruff.Argument{TargetDebateID: &d4IDNull, DebateID: d3.ID, Type: gruff.ARGUMENT_TYPE_CON_IMPACT, Title: "Argument 10", Relevance: 0.2398, Impact: 0.83}
	TESTDB.Create(&a3)
	TESTDB.Create(&a4)
	TESTDB.Create(&a5)
	TESTDB.Create(&a6)
	TESTDB.Create(&a7)
	TESTDB.Create(&a8)
	TESTDB.Create(&a9)
	TESTDB.Create(&a10)

	a3.Debate = &d3
	a4.Debate = &d4
	a5.Debate = &d5
	a6.Debate = &d6
	a7.Debate = &d7
	a8.Debate = &d8
	a9.Debate = &d9
	a10.Debate = &d3

	db := TESTDB
	db = db.Preload("Links")
	db = db.Preload("Contexts")
	db = db.Preload("Values")
	db = db.Preload("Tags")
	db.Where("id = ?", d1.ID).First(&d1)

	d1.ProTruth = []gruff.Argument{a3, a5}
	d1.ConTruth = []gruff.Argument{a4}

	expectedResults, _ := json.Marshal(d1)

	r.GET(fmt.Sprintf("/api/debates/%s", d1.ID.String()))
	res, _ := r.Run(Router())
	assert.Equal(t, string(expectedResults), res.Body.String())
	assert.Equal(t, http.StatusOK, res.Status())
}

func TestCreateDebate(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createDebate()

	r.POST("/api/debates")
	r.SetBody(u1)
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusCreated, res.Status())
}

func TestUpdateDebate(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createDebate()
	TESTDB.Create(&u1)

	r.PUT(fmt.Sprintf("/api/debates/%s", u1.ID))
	r.SetBody(u1)
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusAccepted, res.Status())
}

func TestDeleteDebate(t *testing.T) {
	setup()
	defer teardown()
	r := New(Token)

	u1 := createDebate()
	TESTDB.Create(&u1)

	r.DELETE(fmt.Sprintf("/api/debates/%s", u1.ID))
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
