package main

import (
	"container/heap"
	"fmt"
	"math"
)

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

// order the heap by the cost of the node
func (h Nodes) Less(i, j int) bool { return h[i].cost < h[j].cost }

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

func PrintAssignment(minimum *Node) {
	if minimum.parent == nil {
		return
	}
	PrintAssignment(minimum.parent)
	worker := 'A' + minimum.workerID
	fmt.Printf("Assign worker %c to job %d\n", worker, minimum.jobID)
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

func findMinimumCost(costMatrix []int, gpus int) int {
	jobs := len(costMatrix) / gpus
	h := initializeHeap(Nodes{})

	var assigned []bool
	for i := 0; i < gpus; i++ {
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
	minCost := 0

	for h.Len() > 0 {
		// store the node with the smallest cost
		min := heap.Pop(h).(*Node)
		i := min.workerID + 1

		if i == 4 {
			PrintAssignment(min)
			minCost = min.cost
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
					assigned: min.assigned,
				}
				child.assigned[j] = true
				child.pathCost = min.pathCost + costMatrix[i*gpus+j]
				child.cost = child.pathCost + CalculateCost(costMatrix, i, gpus, jobs, child.assigned)
				heap.Push(h, child)
			}
		}
	}
	return minCost
}

func main() {
	costMatrix := []int{
		9, 2, 7, 8,
		6, 4, 3, 7,
		5, 8, 1, 8,
		7, 6, 9, 4,
	}
	gpus := 4
	optimalCost := findMinimumCost(costMatrix, gpus)
	fmt.Printf("Optimal cost: %d\n", optimalCost)
}
