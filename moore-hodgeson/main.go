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
	remainingJobs []Job
	lastJobID     int
}

type ByDueDate []Job

func (jobs ByDueDate) Len() int {
	return len(jobs)
}

func (jobs ByDueDate) Less(i, j int) bool {
	return jobs[i].dueDate < jobs[j].dueDate
}

func (jobs ByDueDate) Swap(i, j int) {
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
	sort.Sort(ByDueDate(sortedJobs))

	// 2. Make an empty set called `schedulerJobs` and let gamma = 0
	var gamma int = 0

	// 3. Iterate over jobs and append a job accordingly
	for _, job := range sortedJobs {
		if gamma+job.processingTime <= job.dueDate {
			// include the current seen job
			self.scheduledJobs = append(self.scheduledJobs, job)

			// update gamma
			gamma = gamma + job.processingTime
		} else {
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

			// remove the job with the largest processing time from the schedulerJobs
			for _, job := range self.scheduledJobs {
				if job.processingTime == largestProcessingJob.processingTime {
					self.remainingJobs = append(self.remainingJobs, job)
					self.scheduledJobs = append(self.scheduledJobs[:deleteIndex], self.scheduledJobs[deleteIndex+1:]...)
				}
			}

			// update gamma
			gamma = gamma + job.processingTime - largestProcessingJob.processingTime
		}
	}
	return self.scheduledJobs
}

func Equal(job1, job2 Job) bool {
	if job1.dueDate != job2.dueDate && job1.processingTime != job2.processingTime {
		return false
	}
	return true
}

// a function that moves the last element of the duplicated elements in initialJobs into remainingJobs
func (self *JobMaster) MoveDuplicates() {
	for i := 0; i < len(self.initialJobs); i++ {
		for j := i + 1; j < len(self.initialJobs); j++ {
			if Equal(self.initialJobs[i], self.initialJobs[j]) {
				// append the last element of the duplicated elements in initialJobs into remainingJobs
				self.remainingJobs = append(self.remainingJobs, self.initialJobs[j])

				// remove the last element of the duplicated elements in initialJobs
				self.initialJobs = append(self.initialJobs[:j], self.initialJobs[j+1:]...)
			}
		}
	}
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

	// add jobs to the job master
	for _, job := range jobs {
		master.AddJob(job[0], job[1])
	}

	master.MoveDuplicates()
	master.AssignJobs()

	scheduledJobs := master.scheduledJobs
	remainingJobs := master.remainingJobs
	fmt.Println(scheduledJobs)
	fmt.Println(remainingJobs)
}
