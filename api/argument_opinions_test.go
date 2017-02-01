package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/bigokro/gruff-server/gruff"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestListArgumentOpinions(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createArgumentOpinion(TESTDB)
	u2 := createArgumentOpinion(TESTDB)
	TESTDB.Create(&u1)
	TESTDB.Create(&u2)

	expectedResults, _ := json.Marshal([]gruff.ArgumentOpinion{u1, u2})

	r.GET("/api/argument_opinions")
	res, _ := r.Run(Router())
	assert.Equal(t, string(expectedResults), res.Body.String())
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestListArgumentOpinionsPagination(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createArgumentOpinion(TESTDB)
	u2 := createArgumentOpinion(TESTDB)
	TESTDB.Create(&u1)
	TESTDB.Create(&u2)

	r.GET("/api/argument_opinions?start=0&limit=25")
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestGetArgumentOpinions(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createArgumentOpinion(TESTDB)
	TESTDB.Create(&u1)

	expectedResults, _ := json.Marshal(u1)

	r.GET(fmt.Sprintf("/api/argument_opinions/%d", u1.ID))
	res, _ := r.Run(Router())
	assert.Equal(t, string(expectedResults), res.Body.String())
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestCreateArgumentOpinions(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createArgumentOpinion(TESTDB)

	r.POST("/api/argument_opinions")
	r.SetBody(u1)
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusCreated, res.Code)
}

func TestUpdateArgumentOpinions(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createArgumentOpinion(TESTDB)
	TESTDB.Create(&u1)

	r.PUT(fmt.Sprintf("/api/argument_opinions/%d", u1.ID))
	r.SetBody(u1)
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusAccepted, res.Code)
}

func TestDeleteArgumentOpinions(t *testing.T) {
	setup()
	defer teardown()
	r := New(Token)

	u1 := createArgumentOpinion(TESTDB)
	TESTDB.Create(&u1)

	r.DELETE(fmt.Sprintf("/api/argument_opinions/%d", u1.ID))
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusOK, res.Code)
}

func createArgumentOpinion(db *gorm.DB) gruff.ArgumentOpinion {
	d1 := gruff.Claim{Title: "Parent claim", Description: "This is the parent claim"}
	d2 := gruff.Claim{Title: "Child claim", Description: "This is the child claim"}
	db.Create(&d1)
	db.Create(&d2)

	nid := gruff.NullableUUID{d1.ID}
	a := gruff.Argument{
		Title:         "arguments",
		Description:   "arguments",
		Type:          gruff.ARGUMENT_TYPE_PRO_TRUTH,
		TargetClaimID: &nid,
		ClaimID:       d2.ID,
	}
	db.Create(&a)

	u := createUserAO(db)

	ao := gruff.ArgumentOpinion{
		Relevance:  2,
		Impact:     5,
		ArgumentID: a.ID,
		UserID:     uint64(u.ID),
	}

	return ao
}

func createUserAO(db *gorm.DB) gruff.User {
	u := gruff.User{
		Name:     "test",
		Username: "test",
		Email:    "test@test.com",
		Password: "123456",
	}
	password := u.Password
	u.Password = ""
	u.HashedPassword, _ = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return u
}
