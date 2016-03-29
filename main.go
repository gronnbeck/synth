package main

import (
	"log"
	"time"

	"github.com/gronnbeck/synthetic-2/jobs"
	"github.com/gronnbeck/synthetic-2/synth"
)

func main() {

	events := make(chan synth.Event)

	schedule := jobs.Schedule{
		Jobs: []jobs.JobSchedule{
			jobs.JobSchedule{
				Job:         jobs.NewPingURL("https://google.com", "Google.com", events),
				RepeatEvery: 5 * time.Second,
			},
		},
	}

	jobs.RunSchedule(schedule)

	for e := range events {
		log.Println("Oh, Hi!")
		log.Println(e)
	}
}
