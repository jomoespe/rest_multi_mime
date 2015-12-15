package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type Response struct {
	Id   int
	Name string
}

func getResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/xml")

	response := &Response{Id: 1, Name: r.Header.Get("Accept")}

	responseMessage, _ := xml.Marshal(response)
	io.WriteString(w, string(responseMessage))
}

func main() {
	http.HandleFunc("/", getResponse)
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
}
