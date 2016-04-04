package synthesize

import (
	"io/ioutil"

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
	Response Response
}

type Request struct {
	Type string
}

type Response struct {
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
