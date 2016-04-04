package synthesize

import "testing"

func ShouldLoadYAMLJob(t *testing.T) {
	jobYAML := `
name: Test
`
	job, err := loadJobYaml([]byte(jobYAML))

	if err != nil {
		t.Log("Should not fail")
		t.Fail()
	}

	if job.Name != "Test" {
		t.Logf("Expected job name to be %v but was %v", "Test", job.Name)
	}
}
