package synthesize

import "testing"

func Test_Request_Response_with_URL(t *testing.T) {
	server := testTools(200, "")

	action := Action{
		Request:  Request{URL: server.URL},
		Response: ExpectedResponse{StatusCode: 200},
	}

	success, httpResp, err := action.run()

	if err != nil {
		t.Log("Unexpected failure with running the request")
		t.Fatal(err)
	}

	if !success {
		t.Log("Expected actions to succeed but did not")
		t.Logf("Expected status code to be 200 but was %v", httpResp.StatusCode)
		t.Fail()
	}
}

func Test_Request_Response_with_Payload(t *testing.T) {
	server := testTools(200, `{"hello": "world"}`)

	action := Action{
		Request: Request{URL: server.URL},
		Response: ExpectedResponse{
			StatusCode: 200,
			Body:       &map[string]interface{}{"hello": "world"},
		},
	}

	success, _, err := action.run()

	if err != nil {
		t.Log("Unexpected failure with running the request")
		t.Fatal(err)
	}

	if !success {
		t.Log("Expected actions to succeed but did not")
		t.Fail()
	}

}
