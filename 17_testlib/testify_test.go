package testify_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Person struct {
	FirstName, LastName, Phone string
}

func TestSomething(t *testing.T) {
	t.Run("equal", func(t *testing.T) {
		want := 555
		got := 555

		assert.Equal(t, want, got, "they should be equal")
	})

	t.Run("not equal", func(t *testing.T) {
		want := 555
		got := 444

		assert.NotEqual(t, want, got, "they should be equal")
	})

	t.Run("nil", func(t *testing.T) {
		var p *Person

		if assert.Nil(t, p) {
			assert.Nil(t, p)
		}
	})

	t.Run("not nil", func(t *testing.T) {
		pp := &Person{FirstName: "okie"}

		if assert.NotNil(t, pp) {
			assert.Equal(t, "okie", pp.FirstName)
		}
	})
}
