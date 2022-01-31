package main

import (
	"fmt"
	"sort"
)

type Job struct {
	dueDate        int
	processingTime int
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

func MooreHodgson(jobs []Job) []Job {
	// 1. Sort jobs in increasing manner
	sortedJobs := jobs
	sort.Sort(ByDueDate(sortedJobs))

	// 2. Make an empty set called `schedulerJobs` and let gamma = 0
	scheduledJobs := []Job{}
	var gamma int = 0

	for _, job := range sortedJobs {
		if gamma+job.processingTime <= job.processingTime {
			// include the current seen job
			scheduledJobs = append(scheduledJobs, job)

			// update gamma
			gamma = gamma + job.processingTime
		} else {
			// find a job with the largest processing time
			largestProcessingJob := jobs[0]
			for _, job := range scheduledJobs {
				if job.processingTime > largestProcessingJob.processingTime {
					largestProcessingJob = job
				}
			}

			scheduledJobs = append(scheduledJobs, job)

			// remove the job with the largest processing time from the schedulerJobs
			for index, job := range scheduledJobs {
				if job.processingTime == largestProcessingJob.processingTime {
					scheduledJobs = append(scheduledJobs[:index], scheduledJobs[index+1:]...)
				}
			}

			// update gamma
			gamma = gamma + job.processingTime - largestProcessingJob.processingTime
		}
	}

	return scheduledJobs
}

func main() {
	jobs := []Job{
		{6, 4},
		{7, 3},
		{8, 2},
		{9, 5},
		{11, 6},
	}

	scheduledJobs := MooreHodgson(jobs)
	fmt.Println(scheduledJobs)
}
