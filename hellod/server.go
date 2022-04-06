package hellod

import (
	"net/http"

	"github.com/joerx/hellod/hellod/handlers"
)

type Flags struct {
	Address   string
	Unhealthy bool
}

func NewServer(flags Flags) *http.Server {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", handlers.Hello)
	mux.HandleFunc("/health", handlers.Health(flags.Unhealthy))

	s := &http.Server{
		Addr:    flags.Address,
		Handler: mux,
	}

	return s
}
