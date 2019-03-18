package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	addr, ok := os.LookupEnv("ADDRESS")
	if !ok {
		addr = "8080"
	}
	http.HandleFunc("/", handleHello)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func handleHello(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s - %s", req.Method, req.RemoteAddr, req.URL.Path)
	respond200(w, map[string]string{"message": "Hello, world!"})
}

func respond(w http.ResponseWriter, code int, payload interface{}) {
	bytes, _ := json.Marshal(payload)
	w.WriteHeader(code)
	w.Header().Add("Content-type", "application/json")
	w.Write(bytes)
}

func respond200(w http.ResponseWriter, payload interface{}) {
	respond(w, 200, payload)
}

func respond500(w http.ResponseWriter, err error) {
	respond(w, 500, map[string]string{"error": err.Error()})
}
