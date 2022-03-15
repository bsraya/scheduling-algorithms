package main

import (
	"testing"
)

type Test struct {
	numberOfGpus int
	costMatrix   []int
	totalCost    int
	assignments  []Assignment
}

func TestBranchAndBound(t *testing.T) {
	tests := []Test{
		{
			numberOfGpus: 4,
			costMatrix: []int{
				9, 2, 7, 8,
				6, 4, 3, 7,
				5, 8, 1, 8,
				7, 6, 9, 4,
			},
			totalCost: 13,
			assignments: []Assignment{
				{0, 1},
				{1, 0},
				{2, 2},
				{3, 3},
			},
		},
		{
			numberOfGpus: 4,
			costMatrix: []int{
				82, 83, 69, 92,
				77, 37, 49, 92,
				11, 69, 5, 86,
				8, 9, 98, 23,
			},
			totalCost: 140,
			assignments: []Assignment{
				{0, 2},
				{1, 1},
				{2, 0},
				{3, 3},
			},
		},
		{
			numberOfGpus: 3,
			costMatrix: []int{
				2500, 4000, 3500,
				4000, 6000, 3500,
				2000, 4000, 2500,
			},
			totalCost: 9500,
			assignments: []Assignment{
				{0, 1},
				{1, 2},
				{2, 0},
			},
		},
		{
			numberOfGpus: 4,
			costMatrix: []int{
				90, 75, 75, 80,
				30, 85, 55, 65,
				125, 95, 90, 105,
				45, 110, 95, 115,
			},
			totalCost: 275,
			assignments: []Assignment{
				{0, 1},
				{1, 3},
				{2, 2},
				{3, 0},
			},
		},
	}

	for _, test := range tests {
		totalCost, assignments := BranchAndBound(test.costMatrix, test.numberOfGpus)
		if totalCost != test.totalCost {
			t.Errorf("Expected total cost to be %d, but got %d", test.totalCost, totalCost)
		}
		if len(assignments) != len(test.assignments) {
			t.Errorf("Expected %d assignments, but got %d", len(test.assignments), len(assignments))
		}
		// compare assignments with test.assignments to see if they are the same
		for i := range assignments {
			if assignments[i] != test.assignments[i] {
				t.Errorf("Expected GPU %d, but got %d", test.assignments[i], assignments[i])
			}
		}
	}
}
