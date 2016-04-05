package jobs

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gronnbeck/synthetic-2/synth"
)

func testTools(code int, body string) (*httptest.Server, PingURL) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, body)
	}))

	job := NewPingURL(server.URL, "Ping Test", nil)

	return server, job
}

func Test_Ping_Successful(t *testing.T) {
	server, job := testTools(200, "Hello World!")
	defer server.Close()

	events, _ := job.Run()

	resp := events[0]

	if resp.Type != synth.OKType {
		t.Error("Should not return an error when response was 200")
	}
}

func Test_Ping_Failed(t *testing.T) {
	server, job := testTools(404, "Hello World!")
	defer server.Close()

	events, _ := job.Run()

	resp := events[0]

	if resp.Type != synth.ErrorType {
		t.Error("Should return an error when response was 404")
	}

}
