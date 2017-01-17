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
  An Argument connects a Claim to another Claim or Argument
  That is:
     a Claim can be used as an ARGUMENT to either prove or disprove the truth of a claim,
     or to modify the relevance or impact of another argument.

  The TYPE of the argument indicates how the claim (or CLAIM) is being used:
    PRO TRUTH: The Claim is a claim that is being used to prove the truth of another claim
      Ex: "The defendant was in Cincinatti on the date of the murder"
    CON TRUTH: The Claim is used as evidence against another claim
      Ex: "The defendant was hospitalized on the date of the murder"
    PRO RELEVANCE: The Claim is being used to show that another Argument is relevant
      Ex: "The murder occurred in Cincinatti"
    CON RELEVANCE: The Claim is being used to show that another Argument is irrelevant
      Ex: "The murder occurred in the same hospital in which the defendant was hospitalized"
    PRO IMPACT: The Claim is being used to show the importance of another Argument
      Ex: "This argument clearly shows that the defendant has no alibi"
    CON IMPACT: The Claim is being used to diminish the importance of another argument
      Ex: "There is no evidence that the defendant ever left their room"

  A quick explanation of the fields:
    Claim: The Debate (or claim) that is being used as an argument
    Target Claim: The "parent" Claim against which a pro/con truth argument is being made
    Target Argument: In the case of a relevance or impact argument, the argument to which it refers
*/
type Argument struct {
	Identifier
	TargetClaimID    *NullableUUID `json:"targetClaimId,omitempty" sql:"type:uuid"`
	TargetClaim      *Claim        `json:"targetClaim,omitempty"`
	TargetArgumentID *NullableUUID `json:"targetArgId,omitempty" sql:"type:uuid"`
	TargetArgument   *Argument     `json:"targetArg,omitempty"`
	ClaimID          uuid.UUID     `json:"claimId" sql:"type:uuid;not null"`
	Claim            *Claim        `json:"claim,omitempty"`
	Title            string        `json:"title" sql:"not null" valid:"length(3|1000),required"`
	Description      string        `json:"desc" valid:"length(3|4000)"`
	Type             int           `json:"type" sql:"not null"`
	Relevance        float64       `json:"relevance"`
	Impact           float64       `json:"impact"`
	ProRelevance     []Argument    `json:"prorelev,omitempty"`
	ConRelevance     []Argument    `json:"conrelev,omitempty"`
	ProImpact        []Argument    `json:"proimpact,omitempty"`
	ConImpact        []Argument    `json:"conimpact,omitempty"`
}

func (a Argument) ValidateForCreate() GruffError {
	err := a.ValidateField("Title")
	if err != nil {
		return err
	}
	err = a.ValidateField("Description")
	if err != nil {
		return err
	}
	err = a.ValidateField("Type")
	if err != nil {
		return err
	}
	err = a.ValidateIDs()
	if err != nil {
		return err
	}
	err = a.ValidateType()
	if err != nil {
		return err
	}
	return nil
}

func (a Argument) ValidateForUpdate() GruffError {
	return a.ValidateForCreate()
}

func (a Argument) ValidateField(f string) GruffError {
	err := ValidateStructField(a, f)
	return err
}

func (a Argument) ValidateIDs() GruffError {
	if a.ClaimID == uuid.Nil {
		return NewBusinessError("ClaimID: non zero value required;")
	}
	if (a.TargetClaimID == nil || a.TargetClaimID.UUID == uuid.Nil) &&
		(a.TargetArgumentID == nil || a.TargetArgumentID.UUID == uuid.Nil) {
		return NewBusinessError("An Argument must have a target Claim or target Argument ID")
	}
	if a.TargetClaimID != nil && a.TargetArgumentID != nil {
		return NewBusinessError("An Argument can have only one target Claim or target Argument ID")

	}
	return nil
}

func (a Argument) ValidateType() GruffError {
	switch a.Type {
	case ARGUMENT_TYPE_PRO_TRUTH, ARGUMENT_TYPE_CON_TRUTH:
		if a.TargetClaimID == nil || a.TargetClaimID.UUID == uuid.Nil {
			return NewBusinessError("A pro or con truth argument must refer to a target claim")
		}
	case ARGUMENT_TYPE_PRO_RELEVANCE,
		ARGUMENT_TYPE_CON_RELEVANCE,
		ARGUMENT_TYPE_PRO_IMPACT,
		ARGUMENT_TYPE_CON_IMPACT:
		if a.TargetArgumentID == nil || a.TargetArgumentID.UUID == uuid.Nil {
			return NewBusinessError("An impact or relevance argument must refer to a target argument")
		}
	default:
		return NewBusinessError("Type: invalid;")
	}
	return nil
}
