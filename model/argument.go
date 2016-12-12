package model

type Argument struct {
	Model
	ParentID     uint64      `json:"parentId" sql:"not null"`
	Parent       Debate      `json:"parent"`
	DebateID     uint64      `json:"debateId" sql:"not null"`
	Debate       Debate      `json:"debate"`
	Title        string      `json:"title" sql:"not null" valid:"length(3|1000)"`
	Description  string      `json:"desc" valid:"length(3|4000)"`
	Relevance    float64     `json:"relevance"`
	Impact       float64     `json:"impact"`
	ProRelevance []Argument  `json:"prorelev"`
	ConRelevance []Argument  `json:"conrelev"`
	ProImpact    []Argument  `json:"proimpact"`
	ConImpact    []Argument  `json:"conimpact"`
	References   []Reference `json:"refs"`
}
