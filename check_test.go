package pulse

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testCheckError(t *testing.T, statusCode int) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(statusCode), statusCode)
		fmt.Fprintln(w, "Error!")
	}))
	defer ts.Close()

	if err := Check(ts.URL); err == nil {
		t.Errorf("Expecting a %d error", statusCode)
	}
}

func TestCheck(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	if err := Check(ts.URL); err != nil {
		t.Error(err)
	}
}

func TestCheck400(t *testing.T) {
	testCheckError(t, http.StatusBadRequest)
}

func TestCheck401(t *testing.T) {
	testCheckError(t, http.StatusUnauthorized)
}

func TestCheck403(t *testing.T) {
	testCheckError(t, http.StatusForbidden)
}

func TestCheck404(t *testing.T) {
	testCheckError(t, http.StatusNotFound)
}

func TestCheck500(t *testing.T) {
	testCheckError(t, http.StatusInternalServerError)
}

func TestCheck502(t *testing.T) {
	testCheckError(t, http.StatusBadGateway)
}
