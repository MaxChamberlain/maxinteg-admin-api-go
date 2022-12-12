package api

import (
	"fmt"
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
	s.HandleFunc("/user/logout", db.LogoutUser).Methods("GET")
	s.HandleFunc("/user", db.GetUserByToken).Methods("GET")
	s.HandleFunc("/project", db.CreateProject).Methods("POST")
	s.HandleFunc("/project", db.GetProjects).Methods("GET")
	s.HandleFunc("/project/{project_id}", db.GetProject).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:4200", "http://localhost:4200", "https://main--splendorous-cannoli-458d65.netlify.app"},
		AllowCredentials: true,
	})

	handler := c.Handler(s)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	fmt.Println("Starting on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
