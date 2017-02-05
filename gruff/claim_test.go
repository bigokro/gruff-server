package gruff

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateClaim(t *testing.T) {
	setupDB()
	defer teardownDB()

	d := Claim{
		Title:       "The first debate!",
		Description: "A description",
		Truth:       87.55,
	}
	TESTDB.Create(&d)

	assert.True(t, d.ID != ZERO_UUID)
}

func TestUpdateTruth(t *testing.T) {
	setupDB()
	defer teardownDB()

	c1 := Claim{Title: "C1"}
	c2 := Claim{Title: "C2"}
	c3 := Claim{Title: "C3"}
	TESTDB.Create(&c1)
	TESTDB.Create(&c2)
	TESTDB.Create(&c3)

	a1 := Argument{Title: "A1", TargetClaimID: NUUID(c1.ID), ClaimID: c2.ID}
	a2 := Argument{Title: "Heinz 57", TargetClaimID: NUUID(c1.ID), ClaimID: c3.ID}
	TESTDB.Create(&a1)
	TESTDB.Create(&a2)

	c1.UpdateTruth(CTX)
	c2.UpdateTruth(CTX)
	TESTDB.First(&c1)
	TESTDB.First(&c2)
	assert.Equal(t, 0.0, c1.Truth)
	assert.Equal(t, 0.0, c2.Truth)

	co1 := ClaimOpinion{UserID: 1, ClaimID: c1.ID, Truth: 0.5}
	TESTDB.Create(&co1)

	c1.UpdateTruth(CTX)
	c2.UpdateTruth(CTX)
	TESTDB.First(&c1)
	TESTDB.First(&c2)
	assert.Equal(t, 0.5, c1.Truth)
	assert.Equal(t, 0.0, c2.Truth)

	co2 := ClaimOpinion{UserID: 2, ClaimID: c2.ID, Truth: 0.3}
	TESTDB.Create(&co2)

	c1.UpdateTruth(CTX)
	c2.UpdateTruth(CTX)
	TESTDB.First(&c1)
	TESTDB.First(&c2)
	assert.Equal(t, 0.5, c1.Truth)
	assert.Equal(t, 0.3, c2.Truth)

	co3 := ClaimOpinion{UserID: 3, ClaimID: c2.ID, Truth: 0.5}
	TESTDB.Create(&co3)

	c1.UpdateTruth(CTX)
	c2.UpdateTruth(CTX)
	TESTDB.First(&c1)
	TESTDB.First(&c2)
	assert.Equal(t, 0.5, c1.Truth)
	assert.Equal(t, 0.4, c2.Truth)

	co4 := ClaimOpinion{UserID: 4, ClaimID: c1.ID, Truth: 0.9}
	co5 := ClaimOpinion{UserID: 5, ClaimID: c2.ID, Truth: 0.3}
	co6 := ClaimOpinion{UserID: 6, ClaimID: c2.ID, Truth: 0.3}
	co7 := ClaimOpinion{UserID: 7, ClaimID: c2.ID, Truth: 0.9}
	TESTDB.Create(&co4)
	TESTDB.Create(&co5)
	TESTDB.Create(&co6)
	TESTDB.Create(&co7)

	c1.UpdateTruth(CTX)
	c2.UpdateTruth(CTX)
	TESTDB.First(&c1)
	TESTDB.First(&c2)
	assert.Equal(t, 0.7, c1.Truth)
	assert.Equal(t, 0.46, c2.Truth)

	co6.Truth = 0.6
	TESTDB.Save(&co6)

	c1.UpdateTruth(CTX)
	c2.UpdateTruth(CTX)
	TESTDB.First(&c1)
	TESTDB.First(&c2)
	assert.Equal(t, 0.7, c1.Truth)
	assert.Equal(t, 0.52, c2.Truth)
}
