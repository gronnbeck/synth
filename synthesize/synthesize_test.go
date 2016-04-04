package synthesize

import "testing"

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
