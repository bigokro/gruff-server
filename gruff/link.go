package gruff

type Link struct {
	Model
	Identifier
	CreatedByID uint64  `json:"createdById"`
	CreatedBy   *User   `json:"createdBy"`
	Title       string  `json:"title" sql:"not null" valid:"length(3|1000)"`
	Description string  `json:"desc" valid:"length(3|4000)"`
	Url         string  `json:"url" valid:"length(3|4000)"`
	DebateID    *uint64 `json:"debateId"`
	Debate      *Debate `json:"debate,omitempty"`
}
