package gruff

import ()

const CHANGE_TYPE_CREATED_DEBATE int = 1
const CHANGE_TYPE_CREATED_ARGUMENT int = 2
const CHANGE_TYPE_CREATED_DEBATE_AND_ARGUMENT int = 3
const CHANGE_TYPE_MOVED_ARGUMENT int = 11
const CHANGE_TYPE_CLONE_DEBATE int = 21
const CHANGE_TYPE_MERGE_DEBATES int = 31
const CHANGE_TYPE_MERGE_ARGUMENTS int = 32

/*
Types of Changes, and fields used:
- Created Debate: DebateID
- Created Argument: ArgumentID, NewArgType, NewDebateID or NewArgID (parent)
- Created Debate and Argument: ArgumentID, NewArgType, NewDebateID or NewArgID (parent)
- Moved Argument: ArgumentID, OldDebateID or OldArgID, NewDebateID or NewArgID (parent), OldArgType, NewArgType
- Clone Debate:
  - One debate stays
  - New debate created, with same values, context, title and description (must be changed before saving)
  - Arguments stay with main debate
  --> Need Change Type add/remove values and contexts
  --> DebateID, NewDebateID
  --> What about opinions?? I guess it would make a copy
  --> Would there be arguments between old and new debate(s)?
  ------- E.g. Fidel Castro is nice, and ended Apartheid --> Fidel Castro is nice, Fidel Castro ended Apartheid
- Merge Debates:
  - One debate becomes defunct
  - Must have "compatible" values/context (TBD)
  - All arguments attach to "winning" debate
  - Title, description stick with "winning" debate
  - All arguments with "losing" debate as base reattach to "winning" debate
  - Do we need a change log for each of these? Probably...
  - What about opinions? Should also merge...

*/
type ChangeLog struct {
	Model
	UserID      uint64        `json:"userId" sql:"not null"`
	User        *User         `json:"user,omitempty"`
	Type        int           `json:"type" sql:"not null"`
	ArgumentID  *NullableUUID `json:"argumentId,omitempty" sql:"type:uuid"`
	Argument    *Argument     `json:"argument,omitempty"`
	DebateID    *NullableUUID `json:"debateId,omitempty" sql:"type:uuid"`
	Debate      Debate        `json:"debate"`
	OldDebateID *NullableUUID `json:"oldDebateId,omitempty" sql:"type:uuid"`
	OldDebate   *Debate       `json:"oldDebate,omitempty"`
	OldArgID    *NullableUUID `json:"oldArgId,omitempty" sql:"type:uuid"`
	OldArg      *Argument     `json:"oldArg,omitempty"`
	NewDebateID *NullableUUID `json:"newDebateId,omitempty" sql:"type:uuid"`
	NewDebate   *Debate       `json:"newDebate,omitempty"`
	NewArgID    *NullableUUID `json:"newArgId,omitempty" sql:"type:uuid"`
	NewArg      *Argument     `json:"newArg,omitempty"`
	OldArgType  *int          `json:"oldArgType,omitempty"`
	NewArgType  *int          `json:"newArgType,omitempty"`
}
