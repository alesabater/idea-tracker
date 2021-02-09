package data

import (
	"encoding/json"
	"fmt"
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

// Decode functionality

// ToJSON returns ideas as their JSON representation
func (i *Ideas) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

// FromJSON returns a Idea from a JSON format
func (i *Idea) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(i)
}

// RESTful functionlity

// GetIdeas returns a list of ideas
func GetIdeas() Ideas {
	return ideaList
}

// AddIdea adds an idea to the list of ideas
func AddIdea(i *Idea) {
	i.ID = getNextID()
	ideaList = append(ideaList, i)
}

func UpdateIdea(id int, i *Idea) error {
	_, pos, err := findIdea(id)
	if err != nil {
		return err
	}
	i.ID = id
	ideaList[pos] = i
	return nil
}

var ErrIdeaNotFound = fmt.Errorf("Product not found")

func findIdea(id int) (*Idea, int, error) {
	for p, i := range ideaList {
		if id == i.ID {
			return i, p, nil
		}
	}

	return nil, -1, ErrIdeaNotFound
}

// Utility functionality

// getNextID adds an ID
func getNextID() int {
	n := ideaList[len(ideaList)-1]
	return n.ID + 1
}

// IdeaList is a Static list of ideas for testing purposes
var ideaList = []*Idea{
	&Idea{
		ID:          1,
		Name:        "My first idea",
		Description: "My first idea description",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Idea{
		ID:          2,
		Name:        "My second idea",
		Description: "My second idea description",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
