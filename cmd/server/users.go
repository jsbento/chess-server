package server

import (
	"net/http"

	uT "github.com/jsbento/chess-server/cmd/services/users/types"
	"github.com/jsbento/chess-server/pkg/api"
)

func (s *Server) SignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := uT.CreateUserReq{}
		api.ParseAndValidate(r, &req)

		resp, err := s.userS.SignUp(&req)
		api.CheckError(http.StatusInternalServerError, err)
		api.WriteJSON(w, http.StatusOK, resp)
	}
}

func (s *Server) SignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := uT.SignInReq{}
		api.ParseAndValidate(r, &req)

		resp, err := s.userS.SignIn(&req)
		api.CheckError(http.StatusInternalServerError, err)
		api.WriteJSON(w, http.StatusOK, resp)
	}
}
