package responder

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

var errFailedEncoding = []byte(`{"error:"Failed to encode response."}`)

type responder struct {
	http.ResponseWriter
	req *http.Request
}

type Responder interface {
	Success(method string, data interface{})
	Error(method string, err string)
}

type Response struct {
	Success      bool      		`json:"success"`
	Data         interface{} 	`json:"data"`
	Error        *string      	`json:"error"`
}

func New(w http.ResponseWriter, req *http.Request) Responder {
	return &responder{
		ResponseWriter: w,
		req:            req,
	}
}

func (w *responder) writeJSON(body interface{}) {
	bs, err := json.Marshal(body)
	if err != nil {
		log.Println("Something went wrong when encoding JSON output: " + err.Error())
		bs = errFailedEncoding
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	_, err = w.Write(bs)
	if err != nil {
		log.Println("Failed to write HTTP response: " + err.Error())
	}
}

func (w *responder) writeText(message string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if _, err := w.Write([]byte(message)); err != nil {
		log.Println("Failed to write HTTP response: " + err.Error())
	}
}

func (w *responder) write(method string, success bool, data interface{}, err string) {
	switch method {
	case "JSON":
		if (success == true) {
			w.writeJSON(&Response{
				Success: success,
				Data:    data,
				Error: nil,
			})
		} else {
			w.writeJSON(&Response{
				Success: success,
				Error: &err,
				Data:    data,
			})
		}
	}
}

func (w *responder) Success(method string, data interface{}) {
	w.write(strings.ToUpper(method), true, data, "")
}

func (w *responder) Error(method string, err string) {
	w.write(strings.ToUpper(method), false, nil, err)
}