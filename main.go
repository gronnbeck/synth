package main

import (
	"log"
	"time"

	"github.com/gronnbeck/synthetic-2/jobs"
	"github.com/gronnbeck/synthetic-2/synth"
)

func main() {

	events := make(chan synth.Event)
	jobs := []synth.Job{
		jobs.NewPingURL("https://google.com", "Google.com", 5*time.Second, events),
	}

	go func() {
		for _, job := range jobs {
			ticker := time.NewTicker(job.Schedule())
			for range ticker.C {
				log.Println("Hello world!")
				job.Run()
			}
		}
	}()

	for e := range events {
		log.Println("Oh, Hi!")
		log.Println(e)
	}
}
