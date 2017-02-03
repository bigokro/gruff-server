package gruff

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArgumentValidateForCreate(t *testing.T) {
	a := Argument{}

	assert.Equal(t, "Title: non zero value required;", a.ValidateForCreate().Error())

	a.Title = "A"
	assert.Equal(t, "Title: A does not validate as length(3|1000);", a.ValidateForCreate().Error())

	a.Title = "This is a real argument"
	assert.Equal(t, "ClaimID: non zero value required;", a.ValidateForCreate().Error())

	a.Description = "D"
	assert.Equal(t, "Description: D does not validate as length(3|4000);", a.ValidateForCreate().Error())

	a.Description = "This is a real description"
	assert.Equal(t, "ClaimID: non zero value required;", a.ValidateForCreate().Error())

	a.ClaimID = uuid.Nil
	assert.Equal(t, "ClaimID: non zero value required;", a.ValidateForCreate().Error())

	a.ClaimID = uuid.New()
	assert.Equal(t, "An Argument must have a target Claim or target Argument ID", a.ValidateForCreate().Error())

	a.TargetClaimID = &NullableUUID{UUID: uuid.New()}
	assert.Equal(t, "Type: invalid;", a.ValidateForCreate().Error())

	a.Type = 7
	assert.Equal(t, "Type: invalid;", a.ValidateForCreate().Error())

	a.Type = ARGUMENT_TYPE_PRO_TRUTH
	assert.Nil(t, a.ValidateForCreate())

	a.Type = ARGUMENT_TYPE_CON_TRUTH
	assert.Nil(t, a.ValidateForCreate())

	a.Type = ARGUMENT_TYPE_PRO_IMPACT
	assert.Equal(t, "An impact or relevance argument must refer to a target argument", a.ValidateForCreate().Error())

	a.TargetArgumentID = &NullableUUID{UUID: uuid.New()}
	assert.Equal(t, "An Argument can have only one target Claim or target Argument ID", a.ValidateForCreate().Error())

	a.TargetClaimID = nil
	assert.Nil(t, a.ValidateForCreate())

	a.Type = ARGUMENT_TYPE_CON_IMPACT
	assert.Nil(t, a.ValidateForCreate())

	a.Type = ARGUMENT_TYPE_PRO_RELEVANCE
	assert.Nil(t, a.ValidateForCreate())

	a.Type = ARGUMENT_TYPE_CON_RELEVANCE
	assert.Nil(t, a.ValidateForCreate())
}

func TestArgumentValidateForUpdate(t *testing.T) {
	a := Argument{}

	assert.Equal(t, "Title: non zero value required;", a.ValidateForUpdate().Error())

	a.Title = "A"
	assert.Equal(t, "Title: A does not validate as length(3|1000);", a.ValidateForUpdate().Error())

	a.Title = "This is a real argument"
	assert.Equal(t, "ClaimID: non zero value required;", a.ValidateForUpdate().Error())

	a.Description = "D"
	assert.Equal(t, "Description: D does not validate as length(3|4000);", a.ValidateForUpdate().Error())

	a.Description = "This is a real description"
	assert.Equal(t, "ClaimID: non zero value required;", a.ValidateForUpdate().Error())

	a.ClaimID = uuid.Nil
	assert.Equal(t, "ClaimID: non zero value required;", a.ValidateForUpdate().Error())

	a.ClaimID = uuid.New()
	assert.Equal(t, "An Argument must have a target Claim or target Argument ID", a.ValidateForUpdate().Error())

	a.TargetClaimID = &NullableUUID{UUID: uuid.New()}
	assert.Equal(t, "Type: invalid;", a.ValidateForUpdate().Error())

	a.Type = 7
	assert.Equal(t, "Type: invalid;", a.ValidateForUpdate().Error())

	a.Type = ARGUMENT_TYPE_PRO_TRUTH
	assert.Nil(t, a.ValidateForUpdate())

	a.Type = ARGUMENT_TYPE_CON_TRUTH
	assert.Nil(t, a.ValidateForUpdate())

	a.Type = ARGUMENT_TYPE_PRO_IMPACT
	assert.Equal(t, "An impact or relevance argument must refer to a target argument", a.ValidateForUpdate().Error())

	a.TargetArgumentID = &NullableUUID{UUID: uuid.New()}
	assert.Equal(t, "An Argument can have only one target Claim or target Argument ID", a.ValidateForUpdate().Error())

	a.TargetClaimID = nil
	assert.Nil(t, a.ValidateForUpdate())

	a.Type = ARGUMENT_TYPE_CON_IMPACT
	assert.Nil(t, a.ValidateForUpdate())

	a.Type = ARGUMENT_TYPE_PRO_RELEVANCE
	assert.Nil(t, a.ValidateForUpdate())

	a.Type = ARGUMENT_TYPE_CON_RELEVANCE
	assert.Nil(t, a.ValidateForUpdate())
}

func TestOrderByBestArgument(t *testing.T) {
	setupDB()
	defer teardownDB()

	c1 := Claim{Title: "Claim 1", Truth: 0.5}
	c2 := Claim{Title: "Claim 2", Truth: 0.4}
	c3 := Claim{Title: "Claim 3", Truth: 0.6}
	c4 := Claim{Title: "Claim 4", Truth: 0.3}
	c5 := Claim{Title: "Claim 5", Truth: 0.7}
	TESTDB.Create(&c1)
	TESTDB.Create(&c2)
	TESTDB.Create(&c3)
	TESTDB.Create(&c4)
	TESTDB.Create(&c5)

	a1 := Argument{Title: "Argument 1", TargetClaimID: NUUID(c1.ID), ClaimID: c2.ID, Impact: 0.1, Relevance: 0.1}
	a2 := Argument{Title: "Argument 1", TargetClaimID: NUUID(c1.ID), ClaimID: c3.ID, Impact: 0.2, Relevance: 0.0}
	a3 := Argument{Title: "Argument 1", TargetClaimID: NUUID(c1.ID), ClaimID: c4.ID, Impact: 0.6, Relevance: 1.0}
	a4 := Argument{Title: "Argument 1", TargetClaimID: NUUID(c1.ID), ClaimID: c5.ID, Impact: 0.7, Relevance: 0.95}
	TESTDB.Create(&a1)
	TESTDB.Create(&a2)
	TESTDB.Create(&a3)
	TESTDB.Create(&a4)

	args := []Argument{}
	TESTDB.Preload("Claim").Scopes(OrderByBestArgument).Find(&args)
	assert.Equal(t, 4, len(args))
	assert.Equal(t, a4.ID, args[0].ID)
	assert.Equal(t, a3.ID, args[1].ID)
	assert.Equal(t, a1.ID, args[2].ID)
	assert.Equal(t, a2.ID, args[3].ID)
}
