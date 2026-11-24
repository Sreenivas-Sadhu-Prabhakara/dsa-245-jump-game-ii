package main

import "testing"

func TestMinJumps(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
		{[]int{0}, 0},
		{[]int{1, 2}, 1},
		{[]int{3, 2, 1, 0, 4}, -1}, // stranded
		{[]int{1, 1, 1, 1}, 3},
		{[]int{5, 1, 1, 1, 1, 1}, 1},
	}
	for _, c := range cases {
		if got := MinJumps(c.in); got != c.want {
			t.Errorf("MinJumps(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
