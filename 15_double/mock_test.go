package double

import "testing"

type MockSearcher struct {
	phone        string
	methodToCall map[string]bool
}

func (ms *MockSearcher) Search(people []*Person, firstName, lastName string) *Person {
	ms.methodToCall["Search"] = true
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ms.phone,
	}
}

func (ms *MockSearcher) Create(people []*Person, firstName, lastName string) *Person {
	ms.methodToCall["Create"] = true
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ms.phone,
	}
}

func (ms *MockSearcher) ExpectToCall(methodName string) {
	if ms.methodToCall == nil {
		ms.methodToCall = make(map[string]bool)
	}
	ms.methodToCall[methodName] = false
}

func (ms *MockSearcher) Verify(t *testing.T) {
	for methodName, called := range ms.methodToCall {
		if !called {
			t.Errorf("expected to call '%s', but it wasn't.", methodName)
		}
	}
}

func TestCallsWithMocks(t *testing.T) {
	// arange
	fakePhone := "+00 11 22 333"
	phonebook := &Phonebook{}
	mock := &MockSearcher{phone: fakePhone}
	mock.ExpectToCall("Search")

	phone, _ := phonebook.Find(mock, "ok", "bro")

	if phone != fakePhone {
		t.Errorf("want '%s', got '%s'", fakePhone, phone)
	}

	mock.Verify(t)
}
