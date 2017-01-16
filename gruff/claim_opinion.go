package gruff

import (
	"github.com/google/uuid"
)

type ClaimOpinion struct {
	Model
	UserID  uint64    `json:"userId"`
	User    *User     `json:"user,omitempty"`
	ClaimID uuid.UUID `json:"claimId" sql:"type:uuid"`
	Claim   *Claim    `json:"claim,omitempty"`
	Truth   float64   `json:"truth"`
}
