package handlers

import (
	"net/http"

	"github.com/alesabater/idea-tracker/idea-api/data"
)

// swagger:route DELETE /ideas/{id} ideas deleteIdea
// returns a list of ideas
// responses:
//	201: noContent

// DeleteIdea deletes a idea from the database
func (i *IdeaService) DeleteIdea(rw http.ResponseWriter, r *http.Request) {
	id := getIdeaID(r)

	i.l.Println("[DEBUG] deleting record id", id)

	err := data.DeleteIdea(id)
	if err == data.ErrIdeaNotFound {
		i.l.Println("[ERROR] deleting record id does not exist")

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		i.l.Println("[ERROR] deleting record", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)

}
