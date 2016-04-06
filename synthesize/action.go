package synthesize

import (
	"bytes"
	"net/http"
)

type Action struct {
	Request  Request          `yaml:"request"`
	Response ExpectedResponse `yaml:"response"`
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
