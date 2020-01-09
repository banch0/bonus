package main

import (
	"testing"
)

func TestBonus(t *testing.T) {
	sales := []int{15_000, 24_000, 11_000, 8_000, 10_000}
	for i, s := range sales {
		r := bonus(s)
		if i == 0 && r != 750 {
			t.Error("Test error")
		} else if i == 1 && r != 1200 {
			t.Error("Test error")
		} else if i == 2 && r != 550 {
			t.Error("Test error")
		} else if i == 4 && r != 500 {
			t.Error("Test error")
		}
	}
}
