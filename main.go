package main

import (
	"log"
	"time"

	"github.com/gronnbeck/synthetic-2/jobs"
	"github.com/gronnbeck/synthetic-2/synth"
	"github.com/gronnbeck/synthetic-2/synthesize"
)

func main() {

	events := make(chan synth.Event)

	synthesizeJob, _ := synthesize.LoadJobFromFile("./google_ping.yaml")

	schedule := synth.Schedule{
		Jobs: []synth.JobSchedule{
			synth.JobSchedule{
				Job:         jobs.NewPingURL("https://google.com", "Google.com"),
				RepeatEvery: 5 * time.Second,
			},
			synth.JobSchedule{
				Job:         synthesizeJob,
				RepeatEvery: 1 * time.Second,
			},
		},
		EventStream: events,
	}

	synth.RunSchedule(schedule)

	for e := range events {
		log.Println("Oh, Hi!")
		log.Println(e)
	}
}
