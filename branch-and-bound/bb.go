package main

import (
	"container/heap"
	"fmt"
	"math"
)

var assignments []Assignment

type Matrix struct {
	costs        []int
	numberOfGpus int
}

type Assignment struct {
	workerID int
	jobID    int
}

type Node struct {
	// stores the parent of a node
	parent *Node

	// the path cost from the root to the node
	pathCost int

	// the cost of the node
	cost     int
	workerID int
	jobID    int
	assigned []bool
}

type Nodes []*Node

func initializeHeap(nodes []*Node) *Nodes {
	h := Nodes(nodes)
	heap.Init(&h)
	return &h
}

func (h Nodes) Len() int { return len(h) }

// order the heap from the lowest cost all the way up to the highest cost
func (h Nodes) Less(i, j int) bool { return h[i].pathCost > h[j].pathCost }

func (h Nodes) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Nodes) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}

func (h *Nodes) Pop() interface{} {
	// pop the node with the smallest path cost
	old := *h
	n := len(old)
	x := old[n-1]
	old[n-1] = nil // avoid memory leak
	*h = old[0 : n-1]
	return x
}

func AssignJobToNode(minimum *Node) {
	if minimum.parent == nil {
		return
	}
	AssignJobToNode(minimum.parent)
	assignments = append(assignments, Assignment{minimum.workerID, minimum.jobID})
}

func CalculateCost(cost []int, x int, gpus, jobs int, assigned []bool) int {
	totalCost := 0

	available := []bool{}
	for i := 0; i < gpus; i++ {
		available = append(available, true)
	}

	for i := x + 1; i < gpus; i++ {
		min := math.MaxInt
		minIndex := math.MinInt
		for j := 0; j < jobs; j++ {
			if !assigned[j] && available[j] && cost[i*gpus+j] < min {
				minIndex = j
				min = cost[i*gpus+j]
			}
		}
		totalCost += min
		available[minIndex] = false
	}
	return totalCost
}

func BranchAndBound(costMatrix []int, numberOfGpus int) (int, []Assignment) {
	jobs := len(costMatrix) / numberOfGpus
	h := initializeHeap(Nodes{})

	var assigned []bool
	for i := 0; i < numberOfGpus; i++ {
		assigned = append(assigned, false)
	}

	// push a root node into heap
	heap.Push(h, &Node{
		parent:   nil,
		pathCost: 0,
		cost:     0,
		workerID: -1,
		jobID:    -1,
		assigned: assigned,
	})
	cost := 0

	for h.Len() > 0 {
		// store the node with the smallest cost
		min := heap.Pop(h).(*Node)
		i := min.workerID + 1

		if i == numberOfGpus {
			AssignJobToNode(min)
			cost = min.cost
			break
		}
		for j := 0; j < jobs; j++ {
			if !min.assigned[j] {
				child := &Node{
					parent:   min,
					pathCost: 0,
					cost:     0,
					workerID: i,
					jobID:    j,
					assigned: []bool{},
				}
				child.assigned = append(child.assigned, min.assigned...)
				child.assigned[j] = true
				child.pathCost = min.pathCost + costMatrix[i*numberOfGpus+j]
				child.cost = child.pathCost + CalculateCost(costMatrix, i, numberOfGpus, jobs, child.assigned)
				heap.Push(h, child)
			}
		}
	}

	// copy the assignments to a new slice called result
	// and result will be returned
	result := []Assignment{}
	for i := range assignments {
		result = append(result, assignments[i])
	}

	// set global assignment array to empty
	assignments = []Assignment{}

	return cost, result
}

func main() {
	matrices := []Matrix{
		{
			costs: []int{
				9, 2, 7, 8,
				6, 4, 3, 7,
				5, 8, 1, 8,
				7, 6, 9, 4,
			},
			numberOfGpus: 4,
		},
		{
			costs: []int{
				82, 83, 69, 92,
				77, 37, 49, 92,
				11, 69, 5, 86,
				8, 9, 98, 23,
			},
			numberOfGpus: 4,
		},
		{
			costs: []int{
				2500, 4000, 3500,
				4000, 6000, 3500,
				2000, 4000, 2500,
			},
			numberOfGpus: 3,
		},
		{
			costs: []int{
				90, 75, 75, 80,
				30, 85, 55, 65,
				125, 95, 90, 105,
				45, 110, 95, 115,
			},
			numberOfGpus: 4,
		},
	}

	for _, matrix := range matrices {
		optimalCost, assignments := BranchAndBound(matrix.costs, matrix.numberOfGpus)
		fmt.Printf("Optimal cost: %d\n", optimalCost)
		fmt.Printf("Assignments: %v\n", assignments)
	}
}
