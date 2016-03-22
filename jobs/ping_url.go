package jobs

import (
	"fmt"
	"net/http"

	"github.com/gronnbeck/synthetic-2/synth"
)

// PingURL defines a job for pinging an URL.
// It pings and URL and expect the response to have status code 200
type PingURL struct {
	url         string
	application string
	events      chan synth.Event
}

// NewPingURL returns a PingURL struct
func NewPingURL(url string, application string, events chan synth.Event) PingURL {
	return PingURL{
		url:         url,
		application: application,
		events:      events,
	}
}

// Run starts the PingURL job
func (p PingURL) Run() error {
	resp, err := http.Get(p.url)

	if err != nil {
		p.events <- synth.Event{
			ID:          "ping-url-unexpected-error",
			Type:        synth.ErrorType,
			Title:       "Unexpected error",
			Application: p.application,
			Message:     "PingURL on " + p.url + " failed unexpectedly.",
			Tags:        []string{"ping-url"},
		}
		return err
	}

	if resp.StatusCode != 200 {
		p.events <- synth.Event{
			ID:          "ping-url-unexpected-status-code",
			Type:        synth.ErrorType,
			Title:       "Unexpected response code",
			Application: p.application,
			Message: fmt.Sprintf("Expected status code to be 200 but got %v",
				resp.StatusCode),
			Tags: []string{"ping-url", "status-code"},
		}
		return nil
	}

	p.events <- synth.Event{
		ID:          "ping-url-ok",
		Type:        synth.OKType,
		Title:       "Everything went as expected",
		Application: p.application,
		Message:     "Everything is OK. Nothing to report",
		Tags:        []string{"ping-url"},
	}
	return nil
}
