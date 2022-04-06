package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type messageResponse struct {
	Message  string `json:"message"`
	Hostname string `json:"hostname"`
}

type healthCheckResponse struct {
	Status  string `json:"status"`
	Counter int    `json:"counter"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func main() {
	unhealthyFlag := flag.Bool("unhealthy", false, "Make server respond 500 on health check handler")
	flag.Parse()

	log.Printf("Healthy: %t", !*unhealthyFlag)

	addr, ok := os.LookupEnv("ADDRESS")
	if !ok {
		addr = "localhost:8080"
	}

	http.HandleFunc("/", handleHello)
	http.HandleFunc("/health", handleHealth(*unhealthyFlag))

	waitChan := make(chan struct{})
	sigChan := make(chan os.Signal)

	signal.Notify(sigChan, os.Interrupt)

	go func() {
		sig := <-sigChan
		log.Printf("Caught %v, exiting\n", sig)
		close(waitChan)
	}()

	go func() {
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("Server ready at address %s\n", addr)

	<-waitChan

	log.Println("Bye")
}

func handleHello(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s - %s", req.Method, req.RemoteAddr, req.URL.Path)

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	respondOK(w, messageResponse{
		Message:  "Hello World!",
		Hostname: hostname,
	})
}

func handleHealth(unhealthy bool) http.HandlerFunc {
	counter := 0

	return func(w http.ResponseWriter, req *http.Request) {
		counter++
		if unhealthy && counter > 5 {
			respondInternalServerError(w, fmt.Errorf("the server made a boo boo"))
		} else {
			respondOK(w, healthCheckResponse{
				Status:  "OK",
				Counter: counter,
			})
		}
	}
}

func respond(w http.ResponseWriter, code int, payload interface{}) {
	bytes, _ := json.Marshal(payload)
	w.WriteHeader(code)
	w.Header().Add("Content-type", "application/json")
	w.Write(bytes)
}

func respondOK(w http.ResponseWriter, payload interface{}) {
	respond(w, http.StatusOK, payload)
}

func respondInternalServerError(w http.ResponseWriter, err error) {
	respond(w, http.StatusInternalServerError, errorResponse{Error: fmt.Sprintf("%s", err)})
}
