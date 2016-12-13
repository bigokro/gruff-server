package gruff

type Debate struct {
	Model
	UUID        string      `json:"uuid" sql:"not null"`
	CreatedByID uint64      `json:"createdById"`
	CreatedBy   *User       `json:"createdBy"`
	Title       string      `json:"title" sql:"not null" valid:"length(3|1000)"`
	Description string      `json:"desc" valid:"length(3|4000)"`
	Truth       float64     `json:"truth"`
	ProTruth    []Argument  `json:"protruth"`
	ConTruth    []Argument  `json:"contruth"`
	References  []Reference `json:"refs"`
	Contexts    []Context   `json:"contexts"  gorm:"many2many:debate_contexts;"`
	Values      []Value     `json:"values"  gorm:"many2many:debate_values;"`
	Tags        []Tag       `json:"tags"  gorm:"many2many:debate_tags;"`
}
