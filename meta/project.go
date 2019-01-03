package meta

import (
	"encoding/json"
)

// Project ...
type Project struct {
	Name string `json:"name"`
}

// New ...
func New(name string) Project {
	return Project{
		Name: name,
	}
}

func (a Project) String() string {
	b, _ := json.Marshal(a)
	return string(b)
}
