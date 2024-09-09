package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/magrininicolas/ecomgo/service/user"
)

type APIServer struct {
	addr string
	db   *sqlx.DB
}

func NewApiServer(addr string, db *sqlx.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() error {
	r := chi.NewRouter()

	userHandler := user.NewHandler()
	r.Route("/api/v1/users", userHandler.RegisterRoutes)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, r)
}
