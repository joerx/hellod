package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joerx/hellod/hellod/response"
)

type messageResponseRequest struct {
	Method string `json:"method"`
	Path   string `json:"path"`
	Proto  string `json:"proto"`
	Host   string `json:"host"`
	Scheme string `json:"scheme"`
}

type messageResponse struct {
	Message  string                 `json:"message"`
	Hostname string                 `json:"hostname"`
	Request  messageResponseRequest `json:"request"`
}

type healthCheckResponse struct {
	Status  string `json:"status"`
	Counter int    `json:"counter"`
}

func Hello(msg string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Printf("%s %s - %s", req.Method, req.RemoteAddr, req.URL.Path)

		hostname, err := os.Hostname()
		if err != nil {
			log.Fatal(err)
		}

		response.OK(w, messageResponse{
			Message:  msg,
			Hostname: hostname,
			Request: messageResponseRequest{
				Method: req.Method,
				Path:   req.URL.Path,
				Host:   req.URL.Host,
				Scheme: req.URL.Scheme,
			},
		})
	}
}

func Health(unhealthy bool) http.HandlerFunc {
	counter := 0

	return func(w http.ResponseWriter, req *http.Request) {
		counter++
		if unhealthy && counter > 5 {
			response.InternalServerError(w, fmt.Errorf("the server made a boo boo"))
		} else {
			response.OK(w, healthCheckResponse{
				Status:  "OK",
				Counter: counter,
			})
		}
	}
}
