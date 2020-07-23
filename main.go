package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	addr, ok := os.LookupEnv("ADDRESS")
	if !ok {
		addr = "localhost:8080"
	}
	http.HandleFunc("/", handleHello)

	waitChan := make(chan bool)
	sigChan := make(chan os.Signal)

	signal.Notify(sigChan, os.Interrupt)

	go func() {
		sig := <-sigChan
		log.Printf("Caught %v, exiting\n", sig)
		waitChan <- true
	}()

	go func() {
		log.Fatal(http.ListenAndServe(addr, nil))
	}()

	log.Printf("Server ready at address %s\n", addr)

	// wait for closing signal
	<-waitChan
}

func handleHello(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s - %s", req.Method, req.RemoteAddr, req.URL.Path)
	respondOK(w, map[string]string{"message": "Hello, world!"})
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
	respond(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
}

func respondNotFound(w http.ResponseWriter) {
	respond(w, http.StatusNotFound, fmt.Errorf("Not found"))
}
