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
	server.cHub.Run()

	server.r.Get("/play", func(w http.ResponseWriter, r *http.Request) {
		s.ServeChessSocket(server.cHub, w, r)
	})
	return
}
