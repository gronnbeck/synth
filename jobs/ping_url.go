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
}

// NewPingURL returns a PingURL struct
func NewPingURL(url string, application string) PingURL {
	return PingURL{
		url:         url,
		application: application,
	}
}

// Run starts the PingURL job
func (p PingURL) Run() ([]synth.Event, error) {
	resp, err := http.Get(p.url)

	if err != nil {
		errorEvent := synth.Event{
			ID:          "ping-url-unexpected-error",
			Type:        synth.ErrorType,
			Title:       "Unexpected error",
			Application: p.application,
			Message:     "PingURL on " + p.url + " failed unexpectedly.",
			Tags:        []string{"ping-url"},
		}
		return []synth.Event{errorEvent}, err
	}

	if resp.StatusCode != 200 {
		failEvent := synth.Event{
			ID:          "ping-url-unexpected-status-code",
			Type:        synth.ErrorType,
			Title:       "Unexpected response code",
			Application: p.application,
			Message: fmt.Sprintf("Expected status code to be 200 but got %v",
				resp.StatusCode),
			Tags: []string{"ping-url", "status-code"},
		}
		return []synth.Event{failEvent}, nil
	}

	okEvent := synth.Event{
		ID:          "ping-url-ok",
		Type:        synth.OKType,
		Title:       "Everything went as expected",
		Application: p.application,
		Message:     "Everything is OK. Nothing to report",
		Tags:        []string{"ping-url"},
	}
	return []synth.Event{okEvent}, nil
}
