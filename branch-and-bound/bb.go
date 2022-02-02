package main

import (
	"container/heap"
	"fmt"
	"math"
)

const workers = 4

type Node struct {
	parent   *Node
	pathCost int
	cost     int
	workerID int
	jobID    int
	assigned [workers]bool
}

type PriorityQueue []*Node

func print(minimum *Node) {
	if minimum.parent == nil {
		return
	}
	print(minimum.parent)
	fmt.Printf("Assign worker no. %d to do job no. %d\n", minimum.workerID, minimum.jobID)
}

func FindMinimumCost(cost [workers][workers]int) int {
	// Create a priority queue to store live vodes of search tree
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Init(&pq)
	heap.Push(&pq, &Node{0, 0, 0, 0, 0, [workers]bool{false, false, false, false}})

	// while pq is not empty, pop a node from pq
	if node.cost == node.pathCost {
		print(node)
		return
	}
	for i := 0; i < workers; i++ {
		if !node.assigned[i] {
			newNode := &Node{node, node.pathCost, node.pathCost + node.cost[i][node.workerID], i, node.jobID + 1, node.assigned}
			newNode.assigned[i] = true
			heap.Push(&pq, newNode)
		}
	}
	return 0
}

func calculateCost(costMatrix [workers][workers]int, x, y int, assigned [workers]bool) int {
	cost := 0

	available := [workers]bool{true}

	for i := x + 1; i < workers; i++ {
		min := math.MaxInt64
		minIndex := -1
		for j := 0; j < workers; j++ {
			if assigned[j] {
				continue
			}
			if costMatrix[i][j] < min {
				min = costMatrix[i][j]
				minIndex = j
			}
		}
		cost += min
		available[minIndex] = false
	}
	return cost
}

func main() {
	cost := [workers][workers]int{
		{9, 2, 7, 8},
		{6, 4, 3, 7},
		{5, 8, 1, 9},
		{7, 6, 9, 4},
	}
	FindMinimumCost(cost)
}
