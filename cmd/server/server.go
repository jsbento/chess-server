package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	s "github.com/jsbento/chess-server/cmd/server/sockets"
	t "github.com/jsbento/chess-server/cmd/server/types"
	m "github.com/jsbento/chess-server/pkg/mongo"
)

type Server struct {
	r    *chi.Mux
	cHub *s.ChessHub
	m    *m.Store
}

func NewServer() (server *Server, err error) {
	config := &t.ServerConfig{}
	if err := config.LoadAndValidate(); err != nil {
		return nil, err
	}
	store, err := m.NewStore(config.MongoHost, config.MongoDB)
	if err != nil {
		return nil, err
	}

	server = &Server{
		r:    chi.NewRouter(),
		cHub: s.NewChessHub(),
		m:    store,
	}

	server.r.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.Recoverer,
	)

	// ability to http play engine instead of sockets, might be faster
	server.r.Get("/play", func(w http.ResponseWriter, r *http.Request) {
		s.ServeChessSocket(server.cHub, w, r)
	})
	server.r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	return
}

func (s *Server) Start() {
	log.Println("Starting server on port 8080")

	go s.cHub.Run()
	if err := http.ListenAndServe(":8080", s.r); err != nil {
		log.Fatal("Error starting server", err)
	}
}
