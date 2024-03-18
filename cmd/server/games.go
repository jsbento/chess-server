package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	gT "github.com/jsbento/chess-server/cmd/services/games/types"
	"github.com/jsbento/chess-server/pkg/api"
)

func (s *Server) SearchGames() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &gT.SearchGamesReq{}
		api.Parse(r, req)

		games, err := s.gameS.SearchGames(req)
		api.CheckError(http.StatusInternalServerError, err)
		api.WriteJSON(w, http.StatusOK, games)
	}
}

func (s *Server) GetGame() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		game, err := s.gameS.GetGame(chi.URLParam(r, "id"))
		api.CheckError(http.StatusInternalServerError, err)
		api.WriteJSON(w, http.StatusOK, game)
	}
}

func (s *Server) CreateGame() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := gT.Game{}
		api.Parse(r, &req)

		err := s.gameS.CreateGame(&req)
		api.CheckError(http.StatusInternalServerError, err)
		api.WriteJSON(w, http.StatusOK, req)
	}
}

func (s *Server) UpdateGame() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &gT.UpdateGame{}
		api.Parse(r, req)

		req.Id = chi.URLParam(r, "id")

		game, err := s.gameS.UpdateGame(req)
		api.CheckError(http.StatusInternalServerError, err)
		api.WriteJSON(w, http.StatusOK, game)
	}
}

func (s *Server) DeleteGame() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		game, err := s.gameS.DeleteGame(chi.URLParam(r, "id"))
		api.CheckError(http.StatusInternalServerError, err)
		api.WriteJSON(w, http.StatusOK, game)
	}
}
