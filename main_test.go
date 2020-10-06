package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPort(t *testing.T) {
	if port != "8888" {
		t.Errorf("Port is not default value, got: %s, want: %s.", port, "8888")
	}
}

func TestDirectory(t *testing.T) {
	if directory != "." {
		t.Errorf("Directory is not default value, got: %s, want: %s.", directory, ".")
	}
}

func TestFileServer(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.Handler(handleFunc())
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `MiniFileServer.exe`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
