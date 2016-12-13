package gruff

import (
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var ZERO_UUID uuid.UUID

func TestIsIdentifier(t *testing.T) {
	assert.False(t, IsIdentifier(reflect.TypeOf(Tag{})))
	assert.False(t, IsIdentifier(reflect.TypeOf(Value{})))
	assert.False(t, IsIdentifier(reflect.TypeOf(User{})))
	assert.False(t, IsIdentifier(reflect.TypeOf(Context{})))
	assert.False(t, IsIdentifier(reflect.TypeOf(DebateOpinion{})))
	assert.False(t, IsIdentifier(reflect.TypeOf(ArgumentOpinion{})))
	assert.True(t, IsIdentifier(reflect.TypeOf(Debate{})))
	assert.True(t, IsIdentifier(reflect.TypeOf(Argument{})))
	assert.True(t, IsIdentifier(reflect.TypeOf(Link{})))
}

func TestIdentifierGenerateUUID(t *testing.T) {
	d := Debate{}

	assert.Equal(t, ZERO_UUID, d.ID)

	assert.True(t, (&d).GenerateUUID() != ZERO_UUID)
	assert.True(t, d.ID != ZERO_UUID)
}
