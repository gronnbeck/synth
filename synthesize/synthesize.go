package synthesize

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Job struct {
	Name     string
	Schedule Schedule
}

type Schedule struct {
	Duration float32
	Unit     string
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
