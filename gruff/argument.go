package gruff

import (
	"github.com/google/uuid"
)

const ARGUMENT_TYPE_PRO_TRUTH int = 1
const ARGUMENT_TYPE_CON_TRUTH int = 2
const ARGUMENT_TYPE_PRO_RELEVANCE int = 3
const ARGUMENT_TYPE_CON_RELEVANCE int = 4
const ARGUMENT_TYPE_PRO_IMPACT int = 5
const ARGUMENT_TYPE_CON_IMPACT int = 6

type Argument struct {
	Identifier
	ParentID     *NullableUUID `json:"parentId,omitempty" sql:"type:uuid"`
	Parent       *Debate       `json:"parent,omitempty"`
	ArgumentID   *NullableUUID `json:"argumentId,omitempty" sql:"type:uuid"`
	Argument     *Argument     `json:"argument,omitempty"`
	DebateID     uuid.UUID     `json:"debateId" sql:"type:uuid;not null"`
	Debate       Debate        `json:"debate"`
	Title        string        `json:"title" sql:"not null" valid:"length(3|1000)"`
	Description  string        `json:"desc" valid:"length(3|4000)"`
	Type         int           `json:"type" sql:"not null"`
	Relevance    float64       `json:"relevance"`
	Impact       float64       `json:"impact"`
	ProRelevance []Argument    `json:"prorelev"`
	ConRelevance []Argument    `json:"conrelev"`
	ProImpact    []Argument    `json:"proimpact"`
	ConImpact    []Argument    `json:"conimpact"`
}
