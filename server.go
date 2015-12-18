package main

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	"strconv"
	"time"
)

const (
	JSON = "application/json"
	XML  = "text/xml"
)

type Response struct {
	Id   int    `json:"id"   xml:"id,attr"`
	Name string `json:"name" xml:"id"`
}

func getResponse(w http.ResponseWriter, r *http.Request) {

	response := &Response{Id: 1, Name: "Jose"}

	w.Header().Set("Content-Type", "application/xml")
	responseMessage, _ := contentNegotiation(r, *response)
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

func add(a, b int) (int, error) {
	return a + b, nil
}

func contentNegotiation(r *http.Request, response Response) (string, error) {
	var message []byte
	var err error
	str := r.Header.Get("Accept")
	switch str {
	default:
		message = []byte("id=" + strconv.Itoa(response.Id) + ", name=" + response.Name)
	case JSON:
		message, err = json.Marshal(response)
	case XML:
		message, err = xml.Marshal(response)
	}
	return string(message), err
}
