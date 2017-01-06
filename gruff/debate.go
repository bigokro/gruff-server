package gruff

type Debate struct {
	Identifier
	Title       string     `json:"title" sql:"not null" valid:"length(3|1000)"`
	Description string     `json:"desc" valid:"length(3|4000)"`
	Truth       float64    `json:"truth"`
	ProTruth    []Argument `json:"protruth,omitempty"`
	ConTruth    []Argument `json:"contruth,omitempty"`
	Links       []Link     `json:"links,omitempty"`
	Contexts    []Context  `json:"contexts,omitempty"  gorm:"many2many:debate_contexts;"`
	Values      []Value    `json:"values,omitempty"  gorm:"many2many:debate_values;"`
	Tags        []Tag      `json:"tags,omitempty"  gorm:"many2many:debate_tags;"`
}
