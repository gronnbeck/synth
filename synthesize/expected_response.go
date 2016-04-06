package synthesize

import "encoding/json"

type ExpectedResponse struct {
	StatusCode int                     `yaml:"statusCode"`
	Body       *map[string]interface{} `yaml:"body"`
}

func (e ExpectedResponse) TestBody(byt []byte) (bool, error) {
	if e.Body == nil {
		return true, nil
	}

	if string(byt) == "" {
		return leftContains(*e.Body, map[string]interface{}{}), nil
	}

	var actual map[string]interface{}
	err := json.Unmarshal(byt, &actual)
	if err != nil {
		return false, err
	}

	return leftContains(*e.Body, actual), nil
}
