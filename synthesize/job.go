package synthesize

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gronnbeck/synthetic-2/synth"
	"gopkg.in/yaml.v2"
)

type Job struct {
	Name     string   `yaml:"name"`
	Schedule Schedule `yaml:"schedule"`
	Actions  []Action `yaml:"actions"`
}

func (j Job) Run() ([]synth.Event, error) {
	fmt.Println("running job")
	for _, action := range j.Actions {
		passed, resp, err := action.run()

		if err != nil {
			errEvent := synth.Event{
				Application: j.Name,
				Type:        synth.ErrorType,
			}
			return []synth.Event{errEvent}, err
		}

		if !passed {
			failEvent := synth.Event{
				Application: j.Name,
				Type:        "fail",
			}
			log.Println(resp.StatusCode)
			return []synth.Event{failEvent}, nil
		}
	}

	okEvent := synth.Event{
		Application: j.Name,
		Type:        synth.OKType,
	}
	return []synth.Event{okEvent}, nil
}

func LoadJobFromFile(filename string) (*Job, error) {
	byt, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	job, err := LoadJobYAML(byt)
	if err != nil {
		return nil, err
	}

	return job, nil
}

func LoadJobYAML(byt []byte) (*Job, error) {
	job := Job{}
	err := yaml.Unmarshal(byt, &job)
	if err != nil {
		return nil, err
	}

	return &job, nil
}
