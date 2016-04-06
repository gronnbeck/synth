package synthesize

import "net/http"

type Schedule struct {
	Duration float32 `yaml:"duration"`
	Unit     string  `yaml:"unit"`
}

type Request struct {
	Type string `yaml:"type"`
	URL  string `yaml:"url"`
}

func ScheduleJob(job Job) {

}

func (req Request) run() (*http.Response, error) {
	resp, err := http.Get(req.URL)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
