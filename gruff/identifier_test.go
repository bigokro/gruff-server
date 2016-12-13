package gruff

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

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

	assert.Equal(t, "", d.UUID)

	assert.True(t, (&d).GenerateUUID() != "")
	assert.True(t, d.UUID != "")
}
