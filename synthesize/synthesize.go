package synthesize

import (
	"bytes"
	"encoding/json"
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

func (e ExpectedResponse) TestBody(byt []byte) (bool, error) {
	if e.Body == nil {
		return true, nil
	}

	if string(byt) == "" {
		return leftContains(*e.Body, map[string]interface{}{}), nil
	}

	var actual map[string]interface{}
	err := json.Unmarshal(byt, &actual)
	if err != nil {
		return false, err
	}

	return leftContains(*e.Body, actual), nil
}

func (a Action) run() (bool, *http.Response, error) {
	resp, err := a.Request.run()
	if err != nil {
		return false, nil, err
	}

	if resp.StatusCode != a.Response.StatusCode {
		return false, resp, nil
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	resp.Body.Close()

	passBodyTest, err := a.Response.TestBody(buf.Bytes())

	if err != nil {
		return false, resp, err
	}

	if !passBodyTest {
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
		default:
			isLeftContains = isLeftContains && reflect.DeepEqual(v, right[k])
		}
	}
	return isLeftContains
}
