package gruff

import (
	"github.com/satori/go.uuid"
)

type DebateOpinion struct {
	Model
	UserID   uint64    `json:"userId"`
	User     *User     `json:"user,omitempty"`
	DebateID uuid.UUID `json:"debateId"`
	Debate   *Debate   `json:"debate,omitempty"`
	Truth    float64   `json:"truth"`
}
