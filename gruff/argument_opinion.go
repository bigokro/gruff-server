package gruff

import (
	"github.com/google/uuid"
)

type ArgumentOpinion struct {
	Model
	UserID     uint64    `json:"userId"`
	User       *User     `json:"user,omitempty"`
	ArgumentID uuid.UUID `json:"argumentId" sql:"type:uuid"`
	Argument   *Argument `json:"argument,omitempty"`
	Relevance  float64   `json:"relevance"`
	Impact     float64   `json:"impact"`
}
