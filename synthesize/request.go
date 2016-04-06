package synthesize

import "net/http"

type Request struct {
	Type string `yaml:"type"`
	URL  string `yaml:"url"`
}

var client = http.DefaultClient

func (r Request) run() (*http.Response, error) {
	req, err := http.NewRequest(r.Type, r.URL, nil)

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
