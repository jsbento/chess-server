package server

import (
	"net/http"

	cT "github.com/jsbento/chess-server/cmd/services/chess/types"
	"github.com/jsbento/chess-server/pkg/api"
)

func (s *Server) EvalPosition() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := cT.EvalPosReq{}
		api.Parse(r, &req)

		score, err := s.chessS.EvaluatePositon(&req)
		api.CheckError(http.StatusInternalServerError, err)
		api.WriteJSON(w, http.StatusOK, map[string]int{
			"score": score,
		})
	}
}

func (s *Server) SearchPosition() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := cT.SearchPosReq{}
		api.Parse(r, &req)

		bestMove, err := s.chessS.SearchPosition(&req)
		api.CheckError(http.StatusInternalServerError, err)
		api.WriteJSON(w, http.StatusOK, map[string]string{
			"move": bestMove,
		})
	}
}
