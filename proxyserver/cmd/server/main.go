package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Request struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

type Response struct {
	ID      string            `json:"id"`
	Status  int               `json:"status"`
	Headers map[string]string `json:"headers"`
	Length  int               `json:"length"`
}

var requestMap sync.Map

func main() {
	http.HandleFunc("/proxyserver-six.vercel.app", proxyHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if req.Headers == nil {
		http.Error(w, "Header not set", http.StatusBadRequest)
		fmt.Print("Error proxyHandler: Header not set\n")
	}
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	var requestBody io.Reader
	if req.Body != "" {
		requestBody = strings.NewReader(req.Body)
	} else {
		requestBody = nil
	}
	requestID := uuid.New().String()
	httpReq, err := http.NewRequest(req.Method, req.URL, requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for key, value := range req.Headers {
		httpReq.Header.Set(key, value)
	}
	resp, err := client.Do(httpReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	headers := make(map[string]string)
	for key, values := range resp.Header {
		headers[key] = strings.Join(values, ",")
	}
	length, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := Response{
		ID:      requestID,
		Status:  resp.StatusCode,
		Headers: headers,
		Length:  len(length),
	}
	requestMap.Store(requestID, response)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
