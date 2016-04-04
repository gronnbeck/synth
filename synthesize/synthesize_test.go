package synthesize

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var jobYAML = `
name: Test
schedule:
  duration: 5
  unit: seconds
actions:
  - request:
      type: GET
    response:
      statusCode: 200
`

func Test_Should_Load_YAML_Job(t *testing.T) {

	job, err := loadJobYaml([]byte(jobYAML))

	if err != nil {
		t.Log("Parsing should not fail with correct yaml")
		t.Fail()
	}

	if job.Name != "Test" {
		t.Logf("Expected job name to be %v but was %v", "Test", job.Name)
		t.Fail()
	}

	if job.Schedule.Duration != 5 {
		t.Logf("Expected duration to be 5 but was %v", job.Schedule.Duration)
		t.Fail()
	}

	if job.Schedule.Unit != "seconds" {
		t.Logf("Expected unit to be seconds but was %v", job.Schedule.Unit)
		t.Fail()
	}

	if job.Actions[0].Request.Type != "GET" {
		t.Logf("Expected request type to be GET but was %v", job.Actions[0].Request.Type)
	}

	if job.Actions[0].Response.StatusCode != 200 {
		t.Logf("Expected request type to be GET but was %v", job.Actions[0].Response.StatusCode)
	}
}

func Test_Request_Response_with_URL(t *testing.T) {
	server := testTools(200, "")

	action := Action{
		Request:  Request{URL: server.URL},
		Response: Response{StatusCode: 200},
	}

	success, err, httpResp := action.run()

	if err != nil {
		t.Log("Unexpected failure with running the request")
		t.Fail()
	}

	if !success {
		t.Log("Expected actions to succeed but did not")
		t.Logf("Expected status code to be 200 but was %v", httpResp.StatusCode)
		t.Fail()
	}
}

func testTools(code int, body string) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, body)
	}))
	return server
}
