package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	s "github.com/jsbento/chess-server/cmd/server/sockets"
	t "github.com/jsbento/chess-server/cmd/server/types"
	"github.com/jsbento/chess-server/cmd/services/games"
	"github.com/jsbento/chess-server/cmd/services/users"
	"github.com/jsbento/chess-server/pkg/api"
)

type Server struct {
	r     *chi.Mux
	cHub  *s.ChessHub
	gameS *games.GameService
	userS *users.UserService
}

func NewServer() (server *Server, err error) {
	config := &t.ServerConfig{}
	if err := config.LoadAndValidate(); err != nil {
		return nil, err
	}
	gmS, err := games.NewGameService(config)
	if err != nil {
		return nil, err
	}
	usS, err := users.NewUserService(config)
	if err != nil {
		return nil, err
	}

	server = &Server{
		r:     api.NewRouter(),
		cHub:  s.NewChessHub(),
		gameS: gmS,
		userS: usS,
	}

	// ability to http play engine instead of sockets, might be faster
	server.r.Get("/play", func(w http.ResponseWriter, r *http.Request) {
		s.ServeChessSocket(server.cHub, w, r)
	})
	server.r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	// user routes
	server.r.Route("/users", func(r chi.Router) {
		r.Post("/signup", server.SignUp())
		r.Post("/signin", server.SignIn())
	})

	// game routes
	server.r.Route("/games", func(r chi.Router) {
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
