package main

import (
	"testing"
)

func TestBonus(t *testing.T) {
	TestCases := []struct {
		name  string
		sales []int
		want  int
	}{
		{name: "Sale 12000", sales: []int{12_000}, want: 600},
		{name: "Sale 8000", sales: []int{8_000}, want: 0},
		{name: "Sale 10000", sales: []int{10_000}, want: 500},
	}

	for _, testCase := range TestCases {
		got := bonus(testCase.sales)
		if got != testCase.want {
			t.Error("bonus test:", testCase.name, "got:", got, "want:", testCase.want)
		}
	}
}
