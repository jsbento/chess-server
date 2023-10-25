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
