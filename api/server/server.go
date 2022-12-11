package api

import (
	"log"
	"maxinteg-admin-go/api/db"
	jwt "maxinteg-admin-go/helpers/jwt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Server struct {
	*mux.Router
}

func NewServer() *Server {
	db.InitFirebase()
	s := &Server{
		Router: mux.NewRouter(),
	}

	s.Routes()

	return s
}

func (s *Server) Routes() {
	s.Use(jwt.VerifyToken)

	s.HandleFunc("/user/login", db.LoginUser).Methods("POST")
	s.HandleFunc("/user/logout", db.LogoutUser).Methods("POST")
	s.HandleFunc("/user", db.GetUserByToken).Methods("GET")
	s.HandleFunc("/project", db.CreateProject).Methods("PUT")
	s.HandleFunc("/project", db.GetProjects).Methods("GET")
	s.HandleFunc("/project/{project_id}", db.GetProject).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:4200", "http://localhost:4200"},
		AllowCredentials: true,
	})

	handler := c.Handler(s)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handler))
}
