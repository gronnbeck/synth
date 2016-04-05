package synth

import (
	"log"
	"time"
)

// Schedule contains all jobs scheduled
type Schedule struct {
	Jobs        []JobSchedule
	EventStream chan Event
}

// JobSchedule describes how often a job should be repeated
type JobSchedule struct {
	Job         Job
	RepeatEvery time.Duration
}

// RunSchedule runs scheduled job in the background. Will return after all jobs
// has been scheduled.
func RunSchedule(schedule Schedule) {
	for _, scheduledJob := range schedule.Jobs {
		ticker := time.NewTicker(scheduledJob.RepeatEvery)
		go func(job Job, s <-chan time.Time) {
			for range s {
				events, err := job.Run()
				if err != nil {
					log.Printf("Error occured %v", err)
				} else {
					for _, event := range events {
						schedule.EventStream <- event
					}
				}
			}
		}(scheduledJob.Job, ticker.C)
	}
}
