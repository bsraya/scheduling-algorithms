package main

import (
	"fmt"
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
	rejectedJobs  []Job
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
	self.initialJobs = append(self.initialJobs, Job{self.lastJobID, dueDate, processingTime})
	self.lastJobID++
}

// Scheduling algorithm with Moore-Hodgeson
func (self *JobMaster) AssignJobs() []Job {
	// 1. Sort jobs in increasing manner of due date
	sortedJobs := self.initialJobs
	sort.Sort(Jobs(sortedJobs))

	// 2. Let total completion time be 0
	var totalCompletionTime int = 0

	// 3. Iterate over jobs and append a job accordingly
	for _, job := range sortedJobs {
		if totalCompletionTime+job.processingTime <= job.dueDate {
			// include the current seen job
			self.scheduledJobs = append(self.scheduledJobs, job)

			// update totalCompletionTime
			totalCompletionTime = totalCompletionTime + job.processingTime
		} else {
			// include the current seen job
			self.scheduledJobs = append(self.scheduledJobs, job)

			// find a job with the largest processing time
			var largestProcessingJob Job
			var deleteIndex int
			for index, job := range self.scheduledJobs {
				if job.processingTime > largestProcessingJob.processingTime {
					largestProcessingJob = job
					deleteIndex = index
				}
			}

			// remove the job with the largest processing time from schedulerJobs
			// and move it to rejectedJobs
			for _, job := range self.scheduledJobs {
				if job.processingTime == largestProcessingJob.processingTime {
					self.rejectedJobs = append(self.rejectedJobs, job)
					self.scheduledJobs = append(self.scheduledJobs[:deleteIndex], self.scheduledJobs[deleteIndex+1:]...)
				}
			}

			// update totalCompletionTime
			totalCompletionTime = totalCompletionTime + job.processingTime - largestProcessingJob.processingTime
		}
	}
	return self.scheduledJobs
}

func main() {
	var master JobMaster
	jobs := [][]int{
		{6, 4},
		{7, 3},
		{8, 2},
		{9, 5},
		{11, 6},
	}

	// add jobs to the job master
	for _, job := range jobs {
		master.AddJob(job[0], job[1])
	}

	master.AssignJobs()
	scheduledJobs := master.scheduledJobs
	sort.Sort(Jobs(master.rejectedJobs))
	endResult := append(scheduledJobs, master.rejectedJobs...)
	fmt.Println("Scheduled jobs are", master.scheduledJobs)
	fmt.Println("Remaining jobs are", master.rejectedJobs)
	fmt.Println("All the scheduled jobs are", endResult)
}
