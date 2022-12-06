package ticket_test

import (
	"testing"

	"github.com/mallowww/13_testtech/ticket"
)

func TestTicket(t *testing.T) {
	t.Run("Free ticket when age is 3", func(t *testing.T) {
		want := 0.0

		// Blackbox testing
		got := ticket.Price(3)

		if got != want {
			t.Errorf("Price(3) = %f; want %f", got, want)
		}
	})
}
