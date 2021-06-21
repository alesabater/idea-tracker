package data

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

// Idea defines the structure for an API idea
// swagger:model
type Idea struct {
	// the id for this idea
	//
	// required: true
	// min: 0
	ID          int     `json:"id" validate:"required"` // struct tags (go functionality)
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	QAList      IdeaQAs `json:"questions"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Ideas List of ideas
type Ideas []*Idea

func (i *Idea) Validate() error {
	validate := validator.New()
	return validate.Struct(i)
}

// GetIdeas returns a list of ideas
func GetIdeas() Ideas {
	return ideaList
}

// AddIdea adds an idea to the list of ideas
func AddIdea(i *Idea) {
	i.ID = getNextID()
	i.QAList = getDefaultQAList()
	ideaList = append(ideaList, i)
}

// UpdateIdea updates idea
func UpdateIdea(id int, i *Idea) error {
	_, pos, err := findIdea(id)
	if err != nil {
		return err
	}
	i.ID = id
	i.QAList = getDefaultQAList()
	ideaList[pos] = i
	return nil
}

// DeleteProduct deletes a product from the database
func DeleteIdea(id int) error {
	i := findIndexByIdeaID(id)
	if i == -1 {
		return ErrIdeaNotFound
	}

	ideaList = append(ideaList[:i], ideaList[i+1])

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

func findIndexByIdeaID(id int) int {
	for i, p := range ideaList {
		if p.ID == id {
			return i
		}
	}

	return -1
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
		QAList:      getDefaultQAList(),
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Idea{
		ID:          2,
		Name:        "My second idea",
		Description: "My second idea description",
		QAList:      getDefaultQAList(),
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
