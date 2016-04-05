package synthesize

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
		Response: ExpectedResponse{StatusCode: 200},
	}

	success, httpResp, err := action.run()

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

func Test_ExpectedResponse_Comparison(t *testing.T) {
	input := `{"hello": "world", "integer": 4, "float": 3.5}`

	right := map[string]interface{}{}
	err := json.Unmarshal([]byte(input), &right)

	if err != nil {
		t.Fatalf("Should be able to parse json but it failed")
	}

	left := map[string]interface{}{
		"hello":   "world",
		"integer": 4.0,
		"float":   3.5,
	}

	contains := leftContains(left, right)

	if !contains {
		t.Fatal("Left should be contained in right")
	}
}

func Test_ExpectedResponse_Comparison_Complex(t *testing.T) {
	input := `{"complex": {"hello": "world"}}`

	right := map[string]interface{}{}
	err := json.Unmarshal([]byte(input), &right)

	if err != nil {
		t.Fatalf("Should be able to parse json but it failed")
	}

	left := map[string]interface{}{
		"complex": map[string]interface{}{
			"hello": "world",
		},
	}

	contains := leftContains(left, right)

	if !contains {
		t.Fatal("Left should be contained in right")
	}
}

func Test_ExpectedResponse_Comparison_Array(t *testing.T) {
	input := `{"array": [1, 2, 3]}`

	right := map[string]interface{}{}
	err := json.Unmarshal([]byte(input), &right)

	if err != nil {
		t.Fatalf("Should be able to parse json but it failed")
	}

	left := map[string]interface{}{
		"array": []interface{}{1.0, 2.0, 3.0},
	}

	contains := leftContains(left, right)

	if !contains {
		t.Fatal("Left should be contained in right")
	}
}

func Test_ExpectedResponse_Comparison_NotEqual(t *testing.T) {
	input := `{"hello2": "world"}`

	right := map[string]interface{}{}
	err := json.Unmarshal([]byte(input), &right)

	if err != nil {
		t.Fatalf("Should be able to parse json but it failed")
	}

	left := map[string]interface{}{
		"hello": "world",
	}

	contains := leftContains(left, right)

	if contains {
		t.Fatal("Left should not be contained in right")
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

	job, err := loadJobYaml([]byte(spec))

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

func testTools(code int, body string) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, body)
	}))
	return server
}

func Test_leftContains_simple_true(t *testing.T) {
	l := map[string]interface{}{}
	l["test"] = "test"

	r := map[string]interface{}{}
	r["test"] = "test"
	r["ignore"] = "ignore"

	contains := leftContains(l, r)

	if !contains {
		t.Log("Should return true but returned false")
		t.Fail()
	}
}

func Test_leftContains_simple_false(t *testing.T) {
	l := map[string]interface{}{}
	l["tester"] = "test"

	r := map[string]interface{}{}
	r["test"] = "test"
	r["ignore"] = "ignore"

	contains := leftContains(l, r)

	if contains {
		t.Log("Should return false but returned true")
		t.Fail()
	}
}

func Test_leftContains_complex_true(t *testing.T) {
	l := map[string]interface{}{}
	l["test"] = map[string]interface{}{
		"test": "test",
	}

	r := map[string]interface{}{}
	r["test"] = map[string]interface{}{
		"test": "test",
	}
	r["ignore"] = "ignore"

	contains := leftContains(l, r)

	if !contains {
		t.Log("Should return true but returned false")
		t.Fail()
	}
}

func Test_leftContains_complex_false(t *testing.T) {
	l := map[string]interface{}{}
	l["test"] = map[string]interface{}{
		"test": "test",
	}

	r := map[string]interface{}{}
	r["test"] = "test"
	r["ignore"] = "ignore"

	contains := leftContains(l, r)

	if contains {
		t.Log("Should return false but returned true")
		t.Fail()
	}
}

func Test_leftContains_string_array_true(t *testing.T) {
	l := map[string]interface{}{}
	l["test"] = []string{
		"hello",
		"world",
	}

	r := map[string]interface{}{}
	r["test"] = []string{
		"hello",
		"world",
	}
	r["ignore"] = "ignore"

	contains := leftContains(l, r)

	if !contains {
		t.Log("Expected array to work but failed")
		t.Fail()
	}
}

func Test_leftContains_number_array_true(t *testing.T) {
	l := map[string]interface{}{}
	l["test"] = []float64{
		3.3,
		3,
	}

	r := map[string]interface{}{}
	r["test"] = []float64{
		3.3,
		3,
	}
	r["ignore"] = "ignore"

	contains := leftContains(l, r)

	if !contains {
		t.Log("Expected array to work but failed")
		t.Fail()
	}
}

func Test_leftContains_complex_array(t *testing.T) {
	l := map[string]interface{}{}
	l["test"] = []map[string]interface{}{
		map[string]interface{}{
			"test": "test",
		},
	}

	r := map[string]interface{}{}
	r["test"] = []map[string]interface{}{
		map[string]interface{}{
			"test": "test",
		},
	}
	r["ignore"] = "ignore"

	contains := leftContains(l, r)

	if !contains {
		t.Log("Expected complex arrays to work but failed")
		t.Fail()
	}
}
