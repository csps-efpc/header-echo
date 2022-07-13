package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"regexp"
	"testing"
)

func TestIntegration(t *testing.T) {

	cases := []struct {
		pattern        string
		headers        map[string]string
		expectedOutput map[string]string
	}{
		{"^.*$", map[string]string{"X-Shabbadoo": "bingbong"}, map[string]string{"X-Shabbadoo": "bingbong"}},
		{"^.*abba.*$", map[string]string{"X-Shabbadoo": "bingbong"}, map[string]string{"X-Shabbadoo": "bingbong"}},
		{"^.*metallica$", map[string]string{"X-Shabbadoo": "bingbong"}, map[string]string{}},
		{"^.*$", map[string]string{"X-Shabbadoo": "bingbong", "X-Flimflam": "foobarbaz"}, map[string]string{"X-Shabbadoo": "bingbong", "X-Flimflam": "foobarbaz"}},
		{"^.*abba.*$", map[string]string{"X-Shabbadoo": "bingbong", "X-Flimflam": "foobarbaz"}, map[string]string{"X-Shabbadoo": "bingbong"}},
	}

	for _, c := range cases {
		// Set up the proper matcher pattern
		matcherPtr := regexp.MustCompile(c.pattern)
		matcher = *matcherPtr
		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
		// pass 'nil' as the third parameter.
		req, err := http.NewRequest("GET", "/health-check", nil)
		for k, v := range c.headers {
			req.Header.Add(k, v)
		}
		if err != nil {
			t.Fatal(err)
		}
		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(respond)
		handler.ServeHTTP(rr, req)
		// Turn the JSON back to a map
		var bodyMap map[string]string
		t.Logf("Response: %s", rr.Body.String())
		jsonerr := json.Unmarshal(rr.Body.Bytes(), &bodyMap)
		if jsonerr != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(bodyMap, c.expectedOutput) {
			t.Errorf("incorrect output for `%s`: expected `%s` but got `%s`", c.pattern, c.expectedOutput, bodyMap)
		}
	}

}
