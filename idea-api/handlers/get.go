package handlers

import (
	"net/http"

	"github.com/alesabater/idea-tracker/idea-api/data"
)

// swagger:route GET /ideas ideas listIdeas
// returns a list of ideas
// responses:
//	200: ideasResponse

func (i *IdeaService) GetIdeas(rw http.ResponseWriter, w *http.Request) {
	w.Header.Set("Access-Control-Allow-Origin", "*")
	i.l.Println("Handle get Ideas")

	li := data.GetIdeas()
	//fmt.Println(li[0].QAList[0].)

	err := data.ToJSON(li, rw)
	if err != nil {
		http.Error(rw, "Unable to get ideas List", http.StatusInternalServerError)
	}
}
