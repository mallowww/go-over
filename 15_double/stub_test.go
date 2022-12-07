package double

import "testing"

type StubSearcher struct {
	phone string
}

func (ss StubSearcher) Search(people []*Person, firstName, lastName string) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ss.phone,
	}
}

func TestFindReturnsPerson(t *testing.T) {
	fakePhone := "+11 22 333 444"
	phonebook := &Phonebook{}

	ss := StubSearcher{
		phone: fakePhone,
	}
	phone, _ := phonebook.Find(ss, "ok", "lah")

	if phone != fakePhone {
		t.Errorf("want '%s', got '%s'", fakePhone, phone)
	}
}
