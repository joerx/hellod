package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joerx/hellod/hellod/response"
)

type messageResponse struct {
	Message  string `json:"message"`
	Hostname string `json:"hostname"`
}

type healthCheckResponse struct {
	Status  string `json:"status"`
	Counter int    `json:"counter"`
}

func Hello(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s - %s", req.Method, req.RemoteAddr, req.URL.Path)

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	response.OK(w, messageResponse{
		Message:  "Hello World!",
		Hostname: hostname,
	})
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
