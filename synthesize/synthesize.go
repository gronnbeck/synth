package synthesize

import (
	"io/ioutil"
	"net/http"

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

func (a Action) run() (bool, error, *http.Response) {
	resp, err := a.Request.run()
	if err != nil {
		return false, err, nil
	}

	if resp.StatusCode != a.Response.StatusCode {
		return false, nil, resp
	}

	return true, nil, resp
}
