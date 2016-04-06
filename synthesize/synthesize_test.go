package synthesize

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testTools(code int, body string) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, body)
	}))
	return server
}

func Test_Should_Load_YAML_Job(t *testing.T) {

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

	job, err := LoadJobYAML([]byte(jobYAML))

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
		t.Fail()
	}

	if job.Actions[0].Response.StatusCode != 200 {
		t.Logf("Expected request type to be 200 but was %v", job.Actions[0].Response.StatusCode)
		t.Fail()
	}
}

func Test_JobRun_ShouldRunAllActions(t *testing.T) {
	server1 := testTools(202, "")
	server2 := testTools(200, `{"hello": "world"}`)

	job := Job{
		Name: "test",
		Schedule: Schedule{
			Duration: 3.0,
			Unit:     "seconds",
		},
		Actions: []Action{
			Action{
				Request:  Request{URL: server1.URL},
				Response: ExpectedResponse{StatusCode: 202},
			},
			Action{
				Request: Request{URL: server2.URL},
				Response: ExpectedResponse{
					StatusCode: 200,
					Body:       &map[string]interface{}{"hello": "world"},
				},
			},
		},
	}

	_, err := job.Run()

	if err != nil {
		t.Log(err)
		t.Fatal("An error occured with the job")
	}
}

func Test_YAML_ExpectedResponse_Comparison(t *testing.T) {
	spec := `
  name: Test
  schedule:
    duration: 5
    unit: seconds
  actions:
    - request:
        type: GET
      response:
        statusCode: 200
        body:
          hello: world
          world:
            - 1.0
            - 2.0
            - 3.0
  `

	job, err := LoadJobYAML([]byte(spec))

	if err != nil {
		t.Fail()
	}

	input := `{"hello": "world", "world": [1, 2, 3]}`
	actual := map[string]interface{}{}
	err = json.Unmarshal([]byte(input), &actual)

	if err != nil {
		t.Fatal(err)
	}

	exp := job.Actions[0].Response

	contains := leftContains(*exp.Body, actual)

	if !contains {
		t.Fatal("Parsing YAML does not give us the expected request we wanted")
	}

}
