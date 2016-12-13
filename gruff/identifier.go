package gruff

import (
	"github.com/satori/go.uuid"
	"reflect"
	"time"
)

type Identifier struct {
	ID          uuid.UUID  `json:"uuid" sql:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt   time.Time  `json:"-" sql:"DEFAULT:current_timestamp"`
	UpdatedAt   time.Time  `json:"-" sql:"DEFAULT:current_timestamp"`
	DeletedAt   *time.Time `json:"-" settable:"false"`
	CreatedByID uint64     `json:"createdById"`
	CreatedBy   *User      `json:"createdBy"`
}

/*
func (i *Identifier) GenerateUUID() string {
	i.UUID = uuid.NewV4().String()
	return i.UUID
}
*/

func (i *Identifier) GenerateUUID() uuid.UUID {
	i.ID = uuid.NewV4()
	return i.ID
}

func IsIdentifier(t reflect.Type) bool {
	_, is := t.FieldByName("Identifier")
	return is
}
