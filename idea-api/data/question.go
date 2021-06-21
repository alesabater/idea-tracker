package data

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// IdeaQA test
type IdeaQA struct {
	ID       int    `json:"id" validate:"required"`
	Question string `json:"question" validate:"required"`
	Answer   string `json:"answer"`
}

type IdeaQAs []*IdeaQA

func getDefaultQAList() IdeaQAs {
	return listIdeaQuestions
}

func (qa *IdeaQA) Validate() error {
	validator := validator.New()
	return validator.Struct(qa)
}

func (qa *IdeaQA) AnswerQuestion(a string) {
	qa.Answer = a
}

func findQuestion(id int) (*IdeaQA, int, error) {
	for pos, qa := range listIdeaQuestions {
		if qa.ID == id {
			return qa, pos, nil
		}
	}
	return nil, -1, ErrQuestionNotFound
}

var ErrQuestionNotFound = fmt.Errorf("Question not found")

var listIdeaQuestions = []*IdeaQA{
	&IdeaQA{
		ID:       1,
		Question: "What fueled the idea?",
	},
	&IdeaQA{
		ID:       2,
		Question: "Where were you?",
	},
	&IdeaQA{
		ID:       3,
		Question: "Are you passionate about this idea?",
	},
}
