package synthesize

import (
	"bytes"
	"net/http"
)

type Schedule struct {
	Duration float32 `yaml:"duration"`
	Unit     string  `yaml:"unit"`
}

type Action struct {
	Request  Request          `yaml:"request"`
	Response ExpectedResponse `yaml:"response"`
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
