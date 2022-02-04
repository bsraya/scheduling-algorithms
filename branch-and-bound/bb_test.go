package main

import (
	"testing"
)

const n = 4

type Test struct {
	costMatrix [n][n]int
	message    string
}

func BranchAndBoundTest(t *testing.T) {
	tests := []Test{
		Test{
			costMatrix: [n][n]int{
				{9, 2, 7, 8},
				{6, 4, 3, 7},
				{5, 8, 1, 8},
				{7, 6, 9, 4},
			},
			message: "Assign Worker A to Job 1\nAssign Worker B to Job 0\nAssign Worker C to Job 2\nAssign Worker D to Job 3\n13\n",
		},
		Test{
			costMatrix: [n][n]int{
				[n]int{82, 83, 69, 92},
				[n]int{77, 37, 49, 92},
				[n]int{11, 69, 5, 86},
				[n]int{8, 9, 98, 23},
			},
			message: "Assign Worker A to Job 2\nAssign Worker B to Job 1\nAssign Worker C to Job 0\nAssign Worker D to Job 3\n140\n",
		},
		Test{
			costMatrix: [n][n]int{
				[n]int{2500, 4000, 3500},
				[n]int{4000, 6000, 3500},
				[n]int{2000, 4000, 2500},
			},
			message: "Assign Worker A to Job 0\nAssign Worker B to Job 3\nAssign Worker C to Job 2\nAssign Worker D to Job 1\n5000\n",
		},
		Test{
			costMatrix: [n][n]int{
				[n]int{90, 75, 75, 80},
				[n]int{30, 85, 55, 65},
				[n]int{125, 95, 90, 105},
				[n]int{45, 110, 95, 115},
			},
			message: "Assign Worker A to Job 1\nAssign Worker B to Job 3\nAssign Worker C to Job 2\nAssign Worker D to Job 0\n275\n",
		},
	}

	for _, test := range tests {
		if got, want := FindMinimumCost(test.costMatrix), test.message; got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
}
