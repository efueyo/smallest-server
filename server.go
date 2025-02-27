package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"net/url"
)

// Response is strucure that is return for every request
type Response struct {
	OK     bool        `json:"ok"`
	Path   string      `json:"path"`
	Method string      `json:"method"`
	Query  url.Values  `json:"query"`
	Header http.Header `json:"header"`
	Body   string      `json:"body"`
}

// BuildResponse builds the response for a specific request
func BuildResponse(r *http.Request) ([]byte, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	res := Response{
		OK:     true,
		Path:   html.EscapeString(r.URL.Path),
		Method: r.Method,
		Query:  r.URL.Query(),
		Header: r.Header,
		Body:   string(body),
	}
	// ignore error
	return json.Marshal(res)
}

func main() {
	port := 8080
	fmt.Printf("Starting to listen on port: %v\n", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request received: %q\n", html.EscapeString(r.URL.Path))
		res, err := BuildResponse(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %q", err)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(res)
		}
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
