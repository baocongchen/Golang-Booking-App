package main

import (
	"bookings-app/pkg/config"
	"bookings-app/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/majors-suite", handlers.Repo.Majors)
	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Get("/contact", handlers.Repo.Contact)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
