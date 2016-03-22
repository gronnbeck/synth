package synth

import "time"

// Schedule contains all jobs scheduled
type Schedule struct {
	Jobs []JobSchedule
}

// JobSchedule describes how often a job should be repeated
type JobSchedule struct {
	Job         Job
	RepeatEvery time.Duration
}
