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

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	job := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return job
}

func (pq *PriorityQueue) Push(x interface{}) {
	job := x.(*Node)
	*pq = append(*pq, job)
}

func (pq *PriorityQueue) Peek() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	return x
}

func CreateNode(x, y int, assigned [workers]bool, parent *Node) *Node {
	newNode := &Node{
		parent:   parent,
		workerID: x,
		jobID:    y,
		assigned: assigned,
	}
	if parent != nil {
		newNode.assigned[y] = true
	}
	return newNode
}

func FindSmallestElementIndex(pq []*Node) int {
	min := math.MaxInt
	minIndex := -1
	for i := 0; i < len(pq); i++ {
		if pq[i].cost < min {
			min = pq[i].cost
			minIndex = i
		}
	}
	return minIndex
}

func FindMinimumCost(cost [workers][workers]int) int {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	assigned := [workers]bool{false, false, false, false}
	root := CreateNode(-1, -1, assigned, nil)
	root.pathCost = 0
	root.cost = 0
	root.workerID = -1

	pq.Push(root)
	minCost := 0

	for pq.Len() > 0 {
		// find the index of a node with the smallest cost
		minIndex := FindSmallestElementIndex(pq)

		// swap the node with the smallest cost with the last element
		pq.Swap(minIndex, pq.Len()-1)

		// peek the last element that is also the node with the smallest cost
		min := pq.Peek().(*Node)

		// pop the last element
		_ = pq.Pop().(*Node)

		i := min.workerID + 1
		if i == workers {
			// PrintAssignment(min)
			minCost = min.cost
			break
		}

		for j := 0; j < workers; j++ {
			if !min.assigned[j] {
				child := CreateNode(i, j, min.assigned, min)
				child.pathCost = min.pathCost + cost[i][j]
				child.cost = child.pathCost + CalculateCost(cost, i, j, child.assigned)
				pq.Push(child)
			}
		}
	}
	return minCost
}

func PrintAssignment(minimum *Node) {
	if minimum.parent == nil {
		return
	}
	PrintAssignment(minimum.parent)
	worker := 'A' + minimum.workerID
	fmt.Printf("Assign worker %c to job %d\n", worker, minimum.jobID)
}

func CalculateCost(costMatrix [workers][workers]int, x, y int, assigned [workers]bool) int {
	cost := 0

	available := [workers]bool{true, true, true, true}

	for i := x + 1; i < workers; i++ {
		min := math.MaxInt
		minIndex := -1
		for j := 0; j < workers; j++ {
			if !assigned[j] && available[j] && costMatrix[i][j] < min {
				minIndex = j
				min = costMatrix[i][j]
			}
		}
		cost += min
		available[minIndex] = false
	}
	return cost
}

func main() {
	cost := [workers][workers]int{
		{90, 75, 75, 80},
		{30, 85, 55, 65},
		{125, 95, 90, 105},
		{45, 110, 95, 115},
	}
	result := FindMinimumCost(cost)
	fmt.Println(result)
}
