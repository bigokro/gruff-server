package gruff

import (
	"github.com/google/uuid"
)

type Link struct {
	Identifier
	Title       string    `json:"title" sql:"not null" valid:"length(3|1000)"`
	Description string    `json:"desc" valid:"length(3|4000)"`
	Url         string    `json:"url" valid:"length(3|4000)"`
	DebateID    uuid.UUID `json:"debateId" sql:"type:uuid;not null"`
	Debate      *Debate   `json:"debate,omitempty"`
}
