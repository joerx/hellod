package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type messageResponse struct {
	Message string `json:"message"`
}

type errorResponse struct {
	Error string `json:"error"`
}

type counterResponse struct {
	Count int `json:"count"`
}

func main() {
	addr, ok := os.LookupEnv("ADDRESS")
	if !ok {
		addr = "localhost:8080"
	}

	http.HandleFunc("/", handleHello)

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
	respondOK(w, messageResponse{Message: "Hello World!"})
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

func respondNotFound(w http.ResponseWriter) {
	respond(w, http.StatusNotFound, errorResponse{Error: "Not Found"})
}
