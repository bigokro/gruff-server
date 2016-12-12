package model

type Value struct {
	Model
	ParentID    *uint64 `json:"parentId"`
	Parent      *Value  `json:"parent,omitempty"`
	Title       string  `json:"title" sql:"not null" valid:"length(3|1000)"`
	Description string  `json:"desc" valid:"length(3|4000)"`
}
