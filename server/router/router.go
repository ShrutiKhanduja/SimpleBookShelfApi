package router

import (
	"github.com/go-chi/chi"

	server "CODE/server/app"
	"CODE/server/handler"
)

func New(a *server.App) *chi.Mux {
	l := a.Logger()

	r := chi.NewRouter()
	r.Method("GET", "/", handler.NewHandler(a.HandleIndex, l))

	return r
}
