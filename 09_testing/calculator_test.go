package calculator

import (
	"testing"
)

// Test w/ Subtest
func TestSimpleCalculator(t *testing.T) {
	t.Run("should return 3 when input 1, 2", func(t *testing.T) {
		// Arrange
		want := 3

		// Act
		got := sum(1, 2)
		var u int = int(got)

		// Assert
		if u != want {
			t.Errorf("sum(1, 2) = %v; want 3", got)
		}
	})

	t.Run("should return 1 when input 1, 0", func(t *testing.T) {
		got := sum(1, 0)
		if got != 1 {
			t.Errorf("sum(1,0) = %v; want 1", got)
		}
	})

	// writing like this really cozy my eyes
	add := "should return 86.5 when input 44, 42.5"
	t.Run(add, func(t *testing.T) {
		got := sum(44, 42.5)
		if got != 86.5 {
			t.Errorf("sum(44,42.5) = %v; want 86.5", got)
		}
	})

	add_variadic := "should return 70 when input 37, 33"
	t.Run(add_variadic, func(t *testing.T) {
		got := int_sum(37, 33)
		if got != 70 {
			t.Errorf("sum(37,33) = %v; want 70", got)
		}
	})

	sub := "should return 17 when input 10, -7"
	t.Run(sub, func(t *testing.T) {
		got := subtract(10, -7)
		if got != 17 {
			t.Errorf("sub(17,-7) = %v; want 17", got)
		}
	})

	mut := "should return 72 when input 9, 8"
	t.Run(mut, func(t *testing.T) {
		got := multiply(9, 8)
		if got != 72 {
			t.Errorf("sub(9,8) = %v; want 72", got)
		}
	})

	div := "should return 21 when input 84, 4"
	t.Run(div, func(t *testing.T) {
		got := divide(84, 4)
		if got != 21 {
			t.Errorf("sub(84,4) = %v; want 21", got)
		}
	})
}

/* // Test แบบ w/o Subtest
func TestSumRet1when1and0(t *testing.T) {
	got := sum(1, 0)
	if got != 1 {
		t.Errorf("sum(1,0) = %d; want 1", got)
	}
}

func TestSumRetm2whenm1m1(t *testing.T) {
	got := sum(-1, -1)
	if got != -2 {
		t.Errorf("sub(-1,-1) = %d; want -2", got)
	}
}

func TestSubRetMinus10when20m10(t *testing.T) {
	got := subtract(20, 10)
	if got != -10 {
		t.Errorf("sub(20,10) = %v; want -10", got)
	}
}
*/
