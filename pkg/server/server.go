package server

import (
	"encoding/json"
	"exporter-imporoved/pkg/storage"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// TODO graceful shutdown webserver

type Server struct {
	port int
}

// NewServer ...
func NewServer(p int) *Server {
	return &Server{
		port: p,
	}
}

// Start ...
func (s *Server) Start(storage storage.Storable) {
	fmt.Println("Webserver start")
	http.HandleFunc("/health", healthHandler)

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		m, err := storage.GetAll()
		if err != nil {
			// TODO log error
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		c, err := json.Marshal(m)
		if err != nil {
			// TODO log error
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(c)
	})

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(s.port), nil))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO implement healt check here %s!", r.URL.Path[1:])
}
