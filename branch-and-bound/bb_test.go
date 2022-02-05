package main

import (
	"testing"
)

const n = 4

type Test struct {
	costMatrix [n][n]int
	totalCost  int
}

func TestBranchAndBound(t *testing.T) {
	tests := []Test{
		Test{
			costMatrix: [n][n]int{
				{9, 2, 7, 8},
				{6, 4, 3, 7},
				{5, 8, 1, 8},
				{7, 6, 9, 4},
			},
			totalCost: 13,
		},
		Test{
			costMatrix: [n][n]int{
				{82, 83, 69, 92},
				{77, 37, 49, 92},
				{11, 69, 5, 86},
				{8, 9, 98, 23},
			},
			totalCost: 140,
		},
		Test{
			costMatrix: [n][n]int{
				{2500, 4000, 3500},
				{4000, 6000, 3500},
				{2000, 4000, 2500},
			},
			totalCost: 5000,
		},
		Test{
			costMatrix: [n][n]int{
				{90, 75, 75, 80},
				{30, 85, 55, 65},
				{125, 95, 90, 105},
				{45, 110, 95, 115},
			},
			totalCost: 275,
		},
	}

	for _, test := range tests {
		if got, want := FindMinimumCost(test.costMatrix), test.totalCost; got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}
}
