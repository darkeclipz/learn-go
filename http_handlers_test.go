package foohandler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetFoo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handleGetFoo))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	expected := "Foo"
	b, err := io.ReadAll(resp.Body)
	if string(b) != expected {
		t.Errorf("expected %s but we got %s", expected, string(b))
	}
}
