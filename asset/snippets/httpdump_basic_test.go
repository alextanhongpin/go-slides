package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alextanhongpin/core/test/testutil"
)

type EchoRequest struct {
	Message string `json:"message"`
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	var req EchoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func TestEchoHandler(t *testing.T) {
	b := []byte(`{
		"message": "hello world"
	}`)
	r := httptest.NewRequest("POST", "/echo", bytes.NewReader(b))

	testutil.DumpHTTP(t, r, echoHandler)
}
