package main

import "testing"

func TestSum(t *testing.T) {
	total := sum(10, 20)
	want := 30

	if total != want {
		t.Errorf("got %d want %d", total, want)
	}
}

func TestSum2(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, table := range tables {
		total := sum(table.x, table.y)

		if total != table.n {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
		}
	}
}
