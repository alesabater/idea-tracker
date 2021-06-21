package data

import "testing"

func TestIdeaValidation(t *testing.T) {
	i := &Idea{
		ID:   1,
		Name: "My idea",
	}

	err := i.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
