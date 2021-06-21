package data

import "testing"

func TestQuestionValidation(t *testing.T) {
	qa := &IdeaQA{}

	err := qa.Validate()
	if err != nil {
		t.Fatal("Validation failed")
	}
}
