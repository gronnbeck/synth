package synthesize

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Request_post_with_headers(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(200)

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			t.Log(`Authorization header was empty expected "apitoken"`)
			t.Fail()
		} else if authHeader != "apitoken" {
			t.Logf(`Expected "Authorization" to be "apitoken" but was %v`, authHeader)
			t.Fail()
		}

	}))

	r := Request{
		Type: "POST",
		URL:  server.URL,
		Headers: &map[string]string{
			"Authorization": "apitoken",
		},
	}

	r.run()
}
