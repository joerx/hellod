package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/joerx/hellod/hellod"
)

func main() {
	flags := parseFlags()
	s := hellod.NewServer(flags)

	waitChan := make(chan struct{})
	sigChan := make(chan os.Signal)

	signal.Notify(sigChan, os.Interrupt)

	go func() {
		sig := <-sigChan
		log.Printf("caught %v, exiting\n", sig)
		close(waitChan)
	}()

	go s.ListenAndServe()

	log.Printf("server ready at address %s\n", s.Addr)
	<-waitChan
}

func parseFlags() hellod.Flags {
	flags := hellod.Flags{}

	flag.BoolVar(&flags.Unhealthy, "unhealthy", false, "Make server respond 500 on health check handler")
	flag.StringVar(&flags.Address, "address", "localhost:8080", "Address to listen on")
	flag.Parse()

	log.Printf("server flags %#v", flags)
	return flags
}
