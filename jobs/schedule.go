package jobs

import (
	"time"

	"github.com/gronnbeck/synthetic-2/synth"
)

// Schedule contains all jobs scheduled
type Schedule struct {
	Jobs []JobSchedule
}

// JobSchedule describes how often a job should be repeated
type JobSchedule struct {
	Job         synth.Job
	RepeatEvery time.Duration
}
