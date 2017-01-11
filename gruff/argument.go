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

/*
  An Argument connects a Debate to another Debate or Argument
  That is:
     a Debate can be used as an ARGUMENT to either prove or disprove the truth of a debate,
     or to modify the relevance or impact of another argument.

  The TYPE of the argument indicates how the debate (or CLAIM) is being used:
    PRO TRUTH: The Debate is a claim that is being used to prove the truth of another claim
      Ex: "The defendant was in Cincinatti on the date of the murder"
    CON TRUTH: The Debate is used as evidence against another claim
      Ex: "The defendant was hospitalized on the date of the murder"
    PRO RELEVANCE: The Debate is being used to show that another Argument is relevant
      Ex: "The murder occurred in Cincinatti"
    CON RELEVANCE: The Debate is being used to show that another Argument is irrelevant
      Ex: "The murder occurred in the same hospital in which the defendant was hospitalized"
    PRO IMPACT: The Debate is being used to show the importance of another Argument
      Ex: "This argument clearly shows that the defendant has no alibi"
    CON IMPACT: The Debate is being used to diminish the importance of another argument
      Ex: "There is no evidence that the defendant ever left their room"

  A quick explanation of the fields:
    Debate: The Debate (or claim) that is being used as an argument
    Target Debate: The "parent" Debate against which a pro/con truth argument is being made
    Target Argument: In the case of a relevance or impact argument, the argument to which it refers
*/
type Argument struct {
	Identifier
	TargetDebateID   *NullableUUID `json:"targetDebateId,omitempty" sql:"type:uuid"`
	TargetDebate     *Debate       `json:"targetDebate,omitempty"`
	TargetArgumentID *NullableUUID `json:"targetArgId,omitempty" sql:"type:uuid"`
	TargetArgument   *Argument     `json:"targetArg,omitempty"`
	DebateID         uuid.UUID     `json:"debateId" sql:"type:uuid;not null"`
	Debate           *Debate       `json:"debate,omitempty"`
	Title            string        `json:"title" sql:"not null" valid:"length(3|1000)"`
	Description      string        `json:"desc" valid:"length(3|4000)"`
	Type             int           `json:"type" sql:"not null"`
	Relevance        float64       `json:"relevance"`
	Impact           float64       `json:"impact"`
	ProRelevance     []Argument    `json:"prorelev,omitempty"`
	ConRelevance     []Argument    `json:"conrelev,omitempty"`
	ProImpact        []Argument    `json:"proimpact,omitempty"`
	ConImpact        []Argument    `json:"conimpact,omitempty"`
}
