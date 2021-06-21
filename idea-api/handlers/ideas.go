// Package classification of Product API
//
// Documentation for Ideas API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/alesabater/idea-tracker/idea-api/data"
	"github.com/gorilla/mux"
)

// IdeaService struct
type IdeaService struct {
	l *log.Logger
}

// NewIdeaService creates a Ideas handler
func NewIdeaService(l *log.Logger) *IdeaService {
	return &IdeaService{l}
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

func (i IdeaService) AddIdea(rw http.ResponseWriter, r *http.Request) {
	i.l.Println("Handle POST Idea")

	idea := r.Context().Value(KeyIdea{}).(data.Idea)
	data.AddIdea(&idea)
}

func (i IdeaService) UpdateIdea(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Error updating product. Unable to convert ID", http.StatusBadRequest)
	}
	i.l.Println("Handle PUT Idea")

	idea := r.Context().Value(KeyIdea{}).(data.Idea)

	err = data.UpdateIdea(id, &idea)
	if err == data.ErrIdeaNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Error updating product", http.StatusInternalServerError)
	}
}

type KeyIdea struct{}

func (i IdeaService) MiddlewareIdeaValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		idea := data.Idea{}

		err := data.FromJSON(idea, r.Body)
		if err != nil {
			i.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Unable to unmarshal JSON idea", http.StatusBadRequest)
			return
		}

		err = idea.Validate()
		if err != nil {
			i.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, fmt.Sprintf("Unable to unmarshal JSON idea: %s", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyIdea{}, idea)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}

func getIdeaID(r *http.Request) int {
	// parse the product id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}
