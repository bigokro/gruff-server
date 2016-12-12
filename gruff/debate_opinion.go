package gruff

type DebateOpinion struct {
	Model
	UserID   uint64  `json:"userId"`
	User     *User   `json:"user,omitempty"`
	DebateID uint64  `json:"debateId"`
	Debate   *Debate `json:"debate,omitempty"`
	Truth    float64 `json:"truth"`
}
