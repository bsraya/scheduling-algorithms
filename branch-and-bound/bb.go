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

// return the node with the least estimated cost
func (pq PriorityQueue) Top() *Node {
	return pq[0]
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

func FindMinimumCost(cost [workers][workers]int) int {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	assigned := [workers]bool{false}
	root := CreateNode(0, 0, assigned, nil)
	root.pathCost = 0
	root.cost = 0
	root.workerID = 0

	pq.Push(root)
	minCost := 0

	for pq.Len() > 0 {
		min := pq.Pop().(*Node)

		nextWorkerID := min.workerID + 1
		if nextWorkerID == workers {
			PrintAssignment(min)
			minCost = min.cost
			break
		}

		for i := 0; i < workers; i++ {
			if !min.assigned[i] {
				newNode := CreateNode(nextWorkerID, i, min.assigned, min)
				newNode.pathCost = min.pathCost + cost[min.workerID][i]
				newNode.cost = newNode.pathCost + CalculateCost(cost, newNode.workerID, i, newNode.assigned)
				pq.Push(newNode)
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
	char := 'A' + minimum.workerID
	fmt.Printf("Assign worker %s to do job %d\n", string(char), minimum.jobID)
}

func CalculateCost(costMatrix [workers][workers]int, x, y int, assigned [workers]bool) int {
	cost := 0

	available := [workers]bool{true, true, true, true}

	for i := x + 1; i < workers; i++ {
		min := math.MaxInt64
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
		{9, 2, 7, 8},
		{6, 4, 3, 7},
		{5, 8, 1, 9},
		{7, 6, 9, 4},
	}
	result := FindMinimumCost(cost)
	fmt.Println(result)
	// pq := make(PriorityQueue, 0)
	// heap.Init(&pq)
	// assigned := [workers]bool{false}
	// root := CreateNode(0, 0, assigned, nil)
	// root.pathCost = 0
	// root.cost = 0
	// root.workerID = 0

	// pq.Push(root)
	// fmt.Println(pq[0])
}
