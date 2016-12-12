package model

type Tag struct {
	Model
	Title string `json:"title" sql:"not null" valid:"length(3|50)"`
}
