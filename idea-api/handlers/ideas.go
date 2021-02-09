package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/alesabater/idea-tracker/idea-api/data"
)

// IdeaService struct
type IdeaService struct {
	l *log.Logger
}

// NewIdeaService creates a Ideas handler
func NewIdeaService(l *log.Logger) *IdeaService {
	return &IdeaService{l}
}

func (i *IdeaService) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// handle the request for a list of products
	if r.Method == http.MethodGet {
		i.getIdeas(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		i.addIdea(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		i.l.Println("Handle PUT")
		// expect ID in URI
		xp := regexp.MustCompile(`/([0-9]+)`)
		g := xp.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		idStr := g[0][1]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		i.updateIdea(id, rw, r)
		return

	}
	// catch all
	// if no method is satisfied return an error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (i *IdeaService) getIdeas(rw http.ResponseWriter, r *http.Request) {
	i.l.Println("Handle get Ideas")

	li := data.GetIdeas()

	err := li.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to get ideas List", http.StatusInternalServerError)
	}
}

func (i *IdeaService) addIdea(rw http.ResponseWriter, r *http.Request) {
	i.l.Println("Handle POST Idea")

	idea := &data.Idea{}
	err := idea.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal JSON idea", http.StatusBadRequest)
	}

	data.AddIdea(idea)

	i.l.Printf("Idea: %#v", idea)
}

func (i *IdeaService) updateIdea(id int, rw http.ResponseWriter, r *http.Request) {
	i.l.Println("Handle PUT Idea")

	idea := &data.Idea{}
	err := idea.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal JSON idea", http.StatusBadRequest)
	}

	err = data.UpdateIdea(id, idea)
	if err == data.ErrIdeaNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Error updating product", http.StatusInternalServerError)
	}
}
