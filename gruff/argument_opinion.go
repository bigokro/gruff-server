package gruff

import (
	"github.com/satori/go.uuid"
)

type ArgumentOpinion struct {
	Model
	UserID     uint64    `json:"userId"`
	User       *User     `json:"user,omitempty"`
	ArgumentID uuid.UUID `json:"argumentId"`
	Argument   *Argument `json:"argument,omitempty"`
	Relevance  float64   `json:"relevance"`
	Impact     float64   `json:"impact"`
}
