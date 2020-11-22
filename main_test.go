package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)


func TestPostOk(t *testing.T) {
	mcPostBody := map[string]interface{}{
		"timeout": 6,
	}
	body, _ := json.Marshal(mcPostBody)
	r,_ := http.NewRequest("POST", "api/slow",  bytes.NewReader(body))
	w := httptest.NewRecorder()

	returnPost(w,r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, []byte(`{"status":"ok"}`), w.Body.Bytes())
}

func TestPostError(t *testing.T) {
	go startServer()
	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	mcPostBody := map[string]interface{}{
		"timeout": 6000,
	}
	body, _ := json.Marshal(mcPostBody)
	r, _ := http.NewRequest("POST", "http://localhost:8000/api/slow", bytes.NewReader(body))

	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, []byte(`{"error":"timeout too long"}`), body)
}