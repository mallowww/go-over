package ticket

import "testing"

// Boundary test - 0,3/0; 4,15/15; 16,50,/30; 51/5;
// -1
func TestTicketPrice(t *testing.T) {
	t.Run("should return 0 when age = 0", func(t *testing.T) {
		want := 0.0
		age := 0

		got := Price(age)

		if got != want {
			t.Errorf("Price(0) = %f; want %f", got, want)
		}
	})

	t.Run("should return free ticket when age under 3 yrs", func(t *testing.T) {
		want := 0.0
		age := 3

		got := Price(age)

		if got != want {
			t.Errorf("Price(3) = %f; want %f", got, want)
		}
	})

	t.Run("should return 15usd ticket when age at 4 yrs", func(t *testing.T) {
		want := 15.0
		age := 4

		got := Price(age)

		if got != want {
			t.Errorf("Price(4) = %f; want %f", got, want)
		}
	})

	t.Run("should return 15usd ticket when age at 15 yrs", func(t *testing.T) {
		want := 15.0
		age := 15

		got := Price(age)

		if got != want {
			t.Errorf("Price(15) = %f; want %f", got, want)
		}
	})

	t.Run("should return 15usd ticket when age over 15 yrs", func(t *testing.T) {
		want := 30.0
		age := 16

		got := Price(age)

		if got != want {
			t.Errorf("Price(16) = %f; want %f", got, want)
		}
	})

	t.Run("should return 30usd ticket when age at 50 yrs", func(t *testing.T) {
		want := 30.0
		age := 50

		got := Price(age)

		if got != want {
			t.Errorf("Price(50) = %f; want %f", got, want)
		}
	})
	t.Run("should return 30usd ticket when age over 50 yrs", func(t *testing.T) {
		want := 5.0
		age := 51

		got := Price(age)

		if got != want {
			t.Errorf("Price(51) = %f; want %f", got, want)
		}
	})
}
