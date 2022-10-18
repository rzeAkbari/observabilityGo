package server_test

import (
	"github.com/rzeAkbari/observabilityGo/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpServer(t *testing.T) {

	tcs := []struct {
		Method string
		body   string
		result string
		url    string
	}{
		{
			Method: "GET",
			result: "HelloFromGet",
			url:    "products",
		},
		{
			Method: "POST",
			body:   "RAZ",
			result: "HelloFromRaz",
			url:    "product",
		},
		{
			Method: "DELETE",
			result: "HelloFromDeletedID",
			url:    "product/ID",
		},
		{
			Method: "PUT",
			body:   "ID",
			result: "HelloFromPutID",
			url:    "product",
		},
	}
	for _, tc := range tcs {
		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
		// pass 'nil' as the third parameter.
		req, err := http.NewRequest(tc.Method, tc.url, nil)
		if err != nil {
			t.Fatal(err)
		}
		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.ServeProduct)

		handler.ServeHTTP(rr, req)
		want := tc.result
		got := rr.Body.String()

		if got != want {
			t.Errorf("want %s but got %s", want, got)
		}
	}
}
