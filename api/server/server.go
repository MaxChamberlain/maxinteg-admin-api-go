package api

import (
	"maxinteg-admin-go/api/db"

	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}

	s.Routes()
	return s
}

func (s *Server) Routes() {
	// s.HandleFunc("/items", s.createShoppingItem()).Methods(http.MethodPost)
	s.HandleFunc("/user", db.LoginUser).Methods("POST")
}
