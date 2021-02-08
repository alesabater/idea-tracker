package data

import (
	"encoding/json"
	"io"
	"time"
)

// Idea defines the structure for an API idea
type Idea struct {
	ID          int    `json:"id"` // struct tags (go functionality)
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedOn   string `json:"-"`
	UpdatedOn   string `json:"-"`
	DeletedOn   string `json:"-"`
}

// Ideas List of ideas
type Ideas []*Idea

// ToJSON returns ideas as their JSON representation
func (i *Ideas) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

// GetIdeas returns a list of ideas
func GetIdeas() Ideas {
	return IdeaList
}

// IdeaList is a Static list of ideas for testing purposes
var IdeaList = []*Idea{
	&Idea{
		ID:          1,
		Name:        "My first idea",
		Description: "My first idea description",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Idea{
		ID:          1,
		Name:        "My second idea",
		Description: "My second idea description",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
