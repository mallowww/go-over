package ticket

import "testing"

// Boundary test - 0,3/0; 4,15/15; 16,50,/30; 51/5;
// -1

func TestTicketPrice(t *testing.T) {
	// test case
	tests := []struct {
		name string
		age  int
		want float64
	}{ // initiate
		{name: "Free ticket when age is 0", age: 0, want: 0.0},
		{name: "Free ticket when age is 3", age: 3, want: 0.0},
		{name: "15$ ticket when age is 4", age: 4, want: 15.0},
		{name: "15$ ticket when age is 15", age: 15, want: 15.0},
		{"30$ ticket when age is 16", 16, 30.0},
		{"30$ ticket when age is 50", 50, 30.0},
		{"5$ ticket when age is 51", 51, 5.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Price(tt.age)
			if got != tt.want {
				t.Errorf("Price(%d) = %f, want %f", tt.age, got, tt.want)
			}
		})
	}
}
