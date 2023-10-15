package api

import (
	"net/http"

	s "github.com/gorilla/schema"
	j "github.com/helloeave/json"
)

func WriteJSON(w http.ResponseWriter, code int, data interface{}) {
	bytes, err := j.MarshalSafeCollections(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(bytes)
}

func WriteRawJSON(w http.ResponseWriter, code int, data interface{}) {
	bytes, err := j.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(bytes)
}

func Parse(r *http.Request, out interface{}) {
	if r.Method == http.MethodGet {
		if err := s.NewDecoder().Decode(out, r.URL.Query()); err != nil {
			panic(err)
		}
	} else if err := j.NewDecoder(r.Body).Decode(out); err != nil {
		panic(err)
	}
}

type Validator interface {
	Validate() error
}

func ParseAndValidate(r *http.Request, out Validator) {
	Parse(r, out)
	if err := out.Validate(); err != nil {
		panic(err)
	}
}
