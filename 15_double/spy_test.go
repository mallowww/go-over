package double

import "testing"

type SpySearcher struct {
	phone           string
	searchWasCalled bool
	whatIsFirstName string
}

func (ss *SpySearcher) Search(people []*Person, firstName, lastName string) *Person {
	ss.searchWasCalled = true
	ss.whatIsFirstName = firstName
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ss.phone,
	}
}

func TestFindCallsSearchAndReturnsPerson(t *testing.T) {
	fakePhone := "+11 22 333 444"
	phonebook := &Phonebook{}
	spy := &SpySearcher{phone: fakePhone}

	phone, _ := phonebook.Find(spy, "ok", "lah")

	if !spy.searchWasCalled {
		t.Errorf("Expected to call 'Search' in 'Find', but it wasn't.")
	}

	if spy.whatIsFirstName != "ok" {
		t.Error("Expected to call 'Search' with 'ok' as first name, but it wasn't.")
	}

	if phone != fakePhone {
		t.Errorf("want '%s', got '%s'", fakePhone, phone)
	}
}
