package synthesize

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

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

func Test_Action_post_using_correct_http_verb(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Logf("Expected HTTP Method to be post but was %v", r.Method)
			t.Fail()
		}

		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")

		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		r.Body.Close()

		fmt.Fprintln(w, string(buf.Bytes()))
	}))

	action := Action{
		Request: Request{Type: "POST", URL: server.URL},
		Response: ExpectedResponse{
			StatusCode: 200,
			Body:       &map[string]interface{}{"hello": "world"},
		},
	}

	action.run()
}
