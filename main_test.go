package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	NewServer().ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("expected status 200; got %d", res.Code)
	}

	body, _ := io.ReadAll(res.Body)
	if !strings.Contains(string(body), "Welcome to the Go HTTP Server") {
		t.Errorf("unexpected response body: %s", string(body))
	}
}

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	res := httptest.NewRecorder()

	NewServer().ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("expected status 200; got %d", res.Code)
	}

	body, _ := io.ReadAll(res.Body)
	if !strings.Contains(string(body), "ok") {
		t.Errorf("unexpected response body: %s", string(body))
	}
}

// Artificial delay to ensure test lasts ~3 minutes
func TestDelay(t *testing.T) {
	t.Log("Simulating long-running test...")
	time.Sleep(3 * time.Minute)
	t.Log("Completed simulated delay.")
}
