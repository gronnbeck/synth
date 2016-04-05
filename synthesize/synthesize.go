package synthesize

import (
	"io/ioutil"
	"net/http"
	"reflect"

	"gopkg.in/yaml.v2"
)

type Job struct {
	Name     string
	Schedule Schedule
	Actions  []Action
}

type Schedule struct {
	Duration float32
	Unit     string
}

type Action struct {
	Request  Request
	Response ExpectedResponse
}

type Request struct {
	Type string
	URL  string
}

type ExpectedResponse struct {
	StatusCode int
	Body       *map[string]interface{}
}

func loadJobFile(filename string) (*Job, error) {
	byt, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	job, err := loadJobYaml(byt)
	if err != nil {
		return nil, err
	}

	return job, nil
}

func loadJobYaml(byt []byte) (*Job, error) {
	job := Job{}
	err := yaml.Unmarshal(byt, &job)
	if err != nil {
		return nil, err
	}

	return &job, nil
}

func (req Request) run() (*http.Response, error) {
	resp, err := http.Get(req.URL)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a Action) run() (bool, *http.Response, error) {
	resp, err := a.Request.run()
	if err != nil {
		return false, nil, err
	}

	if resp.StatusCode != a.Response.StatusCode {
		return false, resp, nil
	}

	return true, resp, nil
}

func leftContains(left map[string]interface{}, right map[string]interface{}) bool {
	isLeftContains := true
	for k, v := range left {

		if reflect.TypeOf(v) != reflect.TypeOf(right[k]) {
			return false
		}

		switch v.(type) {
		case map[string]interface{}:
			isLeftContains = isLeftContains &&
				leftContains(v.(map[string]interface{}), right[k].(map[string]interface{}))
		case []string:
			isLeftContains = isLeftContains &&
				testEqString(v.([]string), right[k].([]string))
		case []float64:
			isLeftContains = isLeftContains &&
				testEqNumber(v.([]float64), right[k].([]float64))
		case []map[string]interface{}:
			isLeftContains = isLeftContains &&
				testEqComplex(v.([]map[string]interface{}),
					right[k].([]map[string]interface{}))
		default:
			isLeftContains = isLeftContains && v == right[k]
		}
	}
	return isLeftContains
}

func testEqString(a, b []string) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func testEqNumber(a, b []float64) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func testEqComplex(a, b []map[string]interface{}) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i, v1 := range a {
		c := leftContains(v1, b[i])
		if !c {
			return false
		}
	}

	return true
}
