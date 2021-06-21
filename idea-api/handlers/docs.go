package handlers

import "github.com/alesabater/idea-tracker/idea-api/data"

// A list of ideas
// swagger:response ideasResponse
type ideasResponseWrapper struct {
	// All ideas in the system
	// in: body
	Body []data.Idea
}

// No content is returned by this API endpoint
// swagger:response noContent
type noContentResponseWrapper struct {
}

// swagger:parameters deleteIdea
type ideaIDParamsWrapper struct {
	// The id of the idea for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}
