package synthesize

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Request struct {
	Type    string                  `yaml:"type"`
	URL     string                  `yaml:"url"`
	Body    *map[string]interface{} `yaml:"body"`
	Headers *map[string]string      `yaml:"headers"`
}

var client = http.DefaultClient

func (r Request) run() (*http.Response, error) {

	byt, err := json.Marshal(r.Body)

	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(byt)
	req, err := http.NewRequest(r.Type, r.URL, reader)

	if r.Headers != nil {
		for k, v := range *r.Headers {
			req.Header.Set(k, v)
		}
	}

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
