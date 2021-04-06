package servidor

import (
	"net/http"
)

type Server struct {
	Port   string
	Router *Router
}

func NewServer(port string) *Server {
	return &Server{
		Port:   port,
		Router: NewRouter(),
	}
}

func (s *Server) Listen() error {

	http.Handle("/", s.Router)
	err := http.ListenAndServe(s.Port, nil)

	if err != nil {
		return err
	}
	return nil
}
