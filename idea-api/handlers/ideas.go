package handlers

import (
	"log"
	"net/http"

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
}
