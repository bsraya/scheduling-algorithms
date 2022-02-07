package main

import (
	"fmt"
	"math"
	"sort"
)

type Job struct {
	id             int
	dueDate        int
	processingTime int
}

type JobMaster struct {
	initialJobs   []Job
	scheduledJobs []Job
	remainingJobs []Job
	lastJobID     int
}

type Jobs []Job

func (jobs Jobs) Len() int {
	return len(jobs)
}

func (jobs Jobs) Less(i, j int) bool {
	return jobs[i].dueDate < jobs[j].dueDate
}

func (jobs Jobs) Swap(i, j int) {
	jobs[i], jobs[j] = jobs[j], jobs[i]
}

// add job to the job master and increment lastJobID
func (self *JobMaster) AddJob(dueDate int, processingTime int) {
	self.lastJobID++
	self.initialJobs = append(self.initialJobs, Job{self.lastJobID, dueDate, processingTime})
}

// Scheduling algorithm with Moore-Hodgeson
func (self *JobMaster) AssignJobs() []Job {
	// 1. Sort jobs in increasing manner of due date
	sortedJobs := self.initialJobs
	sort.Sort(Jobs(sortedJobs))

	// 2. Make an empty set called `schedulerJobs` and let gamma = 0
	// gamma is the total completion time of schedulued jobs
	var gamma int = 0

	// 3. Iterate over jobs and append a job accordingly
	for _, job := range sortedJobs {
		if gamma+job.processingTime <= job.dueDate {
			// include the current seen job
			self.scheduledJobs = append(self.scheduledJobs, job)

			// update gamma
			gamma += job.processingTime
		} else {
			// include the current seen job
			self.scheduledJobs = append(self.scheduledJobs, job)

			// find a job with the largest processing time
			var threshold = math.MinInt
			// var largestProcessingJob Job
			var deleteIndex int
			for index, job := range self.scheduledJobs {
				if job.processingTime > threshold {
					deleteIndex = index
				}
			}

			// remove the job with the largest processing time from schedulerJobs
			// and move it to remainingJobs
			for _, job := range self.scheduledJobs {
				if job.id == deleteIndex+1 { // +1 because the job id starts from 1
					self.remainingJobs = append(self.remainingJobs, job)
					self.scheduledJobs = append(self.scheduledJobs[:deleteIndex], self.scheduledJobs[deleteIndex+1:]...)
				}
			}
		}
	}
	return self.scheduledJobs
}

func main() {
	// 1. a list of jobs with IDs
	// 2. give each job a unique ID
	// 3. remove duplicate jobs (first come, first serve)
	// 3. pass the list of jobs to the job master
	// 4. the job master schedules jobs so that we have the least amount of late jobs
	// 5. append the remaining jobs to the scheduled jobs

	var master JobMaster
	jobs := [][]int{
		{6, 4},
		{7, 3},
		{8, 2},
		{9, 5},
		{11, 6},
	}

	// jobs := [][]int{
	// 	{6, 4},
	// 	{8, 1},
	// 	{9, 6},
	// 	{11, 3},
	// 	{20, 6},
	// 	{25, 8},
	// 	{28, 7},
	// 	{35, 10},
	// }

	// add jobs to the job master
	for _, job := range jobs {
		master.AddJob(job[0], job[1])
	}

	master.AssignJobs()

	scheduledJobs := master.scheduledJobs
	sort.Sort(Jobs(master.remainingJobs))

	// append the remaining jobs to the scheduled jobs
	endResult := append(scheduledJobs, master.remainingJobs...)

	fmt.Println("Scheduled jobs are", master.scheduledJobs)
	fmt.Println("Remaining jobs are", master.remainingJobs)
	fmt.Println("All the scheduled jobs are", endResult)
}
