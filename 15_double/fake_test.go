package double

import "testing"

type FakeResearcher struct {
}

func (fs FakeResearcher) Search(people []*Person, firstName, lastName string) *Person {
	if len(people) == 0 {
		return nil
	}
	return people[0]
}

// return empty string for no person
func TestFindCallsSearch(t *testing.T) {
	phonebook := &Phonebook{}
	fake := &FakeResearcher{}

	phone, _ := phonebook.Find(fake, "ok", "lah")

	if phone != "" {
		t.Errorf("wanted '', got '%s'", phone)
	}
}
