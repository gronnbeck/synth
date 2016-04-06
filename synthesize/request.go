package synthesize

import "net/http"

type Request struct {
	Type string `yaml:"type"`
	URL  string `yaml:"url"`
}

func (req Request) run() (*http.Response, error) {
	resp, err := http.Get(req.URL)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
