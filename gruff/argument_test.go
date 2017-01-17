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
