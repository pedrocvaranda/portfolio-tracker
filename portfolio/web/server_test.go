package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	server := MustNewServer()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	server.Handler().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	if !bytes.Contains(rec.Body.Bytes(), []byte("Portfolio Tracker")) {
		t.Fatalf("index response did not contain app title")
	}
}

func TestSimulate(t *testing.T) {
	server := MustNewServer()
	body := bytes.NewBufferString(`{"ticker":"AAPL","startPrice":100,"drift":0.08,"volatility":0.22,"days":10,"paths":20,"seed":42}`)
	req := httptest.NewRequest(http.MethodPost, "/api/simulate", body)
	rec := httptest.NewRecorder()

	server.Handler().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d: %s", rec.Code, http.StatusOK, rec.Body.String())
	}
	var resp simulateResponse
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatal(err)
	}
	if resp.Ticker != "AAPL" || resp.Count != 20 || len(resp.Paths) == 0 {
		t.Fatalf("unexpected response: %+v", resp)
	}
}
