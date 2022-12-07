package double

import "testing"

type DummySearcher struct {
}

func (ds DummySearcher) Search(people []*Person, firstName, lastName string) *Person {
	return &Person{}
}

// when firstName or lastName empty
func TestFindReturnsError(t *testing.T) {
	phonebook := &Phonebook{}
	want := ErrMissingArgs

	dd := DummySearcher{}
	_, got := phonebook.Find(dd, "", "")

	if got != want {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}
