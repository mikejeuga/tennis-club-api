package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
}

func NewServer() *http.Server {
	router := mux.NewRouter()
	s := Server{}

	router.HandleFunc("/", s.HealthCheck)

	return &http.Server{
		Addr:    ":8072",
		Handler: router,
	}
}

func (s Server) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Tennis server is up!")
}
