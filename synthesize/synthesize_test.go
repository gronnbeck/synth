package synthesize

import "testing"

func Test_Should_Load_YAML_Job(t *testing.T) {
	jobYAML := `
name: Test
schedule:
  duration: 5
  unit: seconds
`
	job, err := loadJobYaml([]byte(jobYAML))

	if err != nil {
		t.Log("Should not fail")
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
}
