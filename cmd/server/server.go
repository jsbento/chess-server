package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	s "github.com/jsbento/chess-server/cmd/server/sockets"
)

type Server struct {
	r    *chi.Mux
	cHub *s.ChessHub
}

func NewServer() (server *Server) {
	server = &Server{
		r:    chi.NewRouter(),
		cHub: s.NewChessHub(),
	}

	server.r.Get("/play", func(w http.ResponseWriter, r *http.Request) {
		s.ServeChessSocket(server.cHub, w, r)
	})
	return
}

func (s *Server) Start() {
	go s.cHub.Run()
	if err := http.ListenAndServe(":8080", s.r); err != nil {
		panic(err)
	}
}
