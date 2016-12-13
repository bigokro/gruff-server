package gruff

import (
	"github.com/satori/go.uuid"
	"reflect"
)

type Identifier struct {
	UUID string `json:"uuid" sql:"not null"`
}

func (i *Identifier) GenerateUUID() string {
	i.UUID = uuid.NewV4().String()
	return i.UUID
}

func IsIdentifier(t reflect.Type) bool {
	_, is := t.FieldByName("Identifier")
	return is
}
