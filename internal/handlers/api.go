package handlers

import (
	"github.com/Abhishekkapoor98/goapi/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {

	//Global Middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		// Middleware for /account routes
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}
