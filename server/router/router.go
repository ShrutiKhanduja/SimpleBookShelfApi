package router

import (
	"github.com/go-chi/chi"

	server "CODE/server/app"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.MethodFunc("GET", "/", server.HandleIndex)

	return r
}
