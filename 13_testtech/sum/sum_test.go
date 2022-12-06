package sum

import "testing"

func TestSum(t *testing.T) {
	got := sum([]int{}...)
	want := 0
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

// func TestSum(t *testing.T) {
// 	want := 5

// 	got := sum(2, 3)

// 	if got != want {
// 		t.Error("Expected", 5, "Got", got)
// 	}
// }

func TestSumAll(t *testing.T) {
	want := 7
	xs := []int{2, 3, 3, -1}

	// got := sum([]int{2, 3, 3, -1}...)
	got := sum(xs...)

	if got != want {
		t.Error("Expected", 8, "Got", got)
	}
}
