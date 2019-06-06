package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPricription(t *testing.T) {
	var jsonStr = []byte(`{"alertname": "test"}`)
	req, err := http.NewRequest("POST", "/script", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(scriptHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code, "OK response is expected")
	expected := "failed"
	if rr.Body.String() != expected {
		t.Errorf("alert send failed")
	}
	log.Println(rr.Body)
}
