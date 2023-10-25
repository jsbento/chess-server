package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	s "github.com/jsbento/chess-server/cmd/server/sockets"
	t "github.com/jsbento/chess-server/cmd/server/types"
	"github.com/jsbento/chess-server/cmd/services/chess"
	"github.com/jsbento/chess-server/cmd/services/games"
	"github.com/jsbento/chess-server/cmd/services/users"
	"github.com/jsbento/chess-server/pkg/api"
	"github.com/jsbento/chess-server/pkg/auth"
)

type Server struct {
	r      *chi.Mux
	cHub   *s.ChessHub
	gameS  *games.GameService
	userS  *users.UserService
	chessS *chess.ChessService
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
	chS, err := chess.NewChessService(config)
	if err != nil {
		return nil, err
	}

	server = &Server{
		r:      api.NewRouter(),
		cHub:   s.NewChessHub(),
		gameS:  gmS,
		userS:  usS,
		chessS: chS,
	}

	server.r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler)

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
		r.Get("/", auth.CheckAuth(server.SearchGames()))
		r.Get("/{id}", auth.CheckAuth(server.GetGame()))
	})

	// chess routes (for engine)
	server.r.Route("/chess", func(r chi.Router) {
		r.Post("/eval", server.EvalPosition())
		r.Post("/search", server.SearchPosition())
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
