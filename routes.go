package main

import (
	"net/http"

	"github.com/TYHXX/go-miniBookingSystem/pkg/config"
	handerls "github.com/TYHXX/go-miniBookingSystem/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handerls.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handerls.Repo.About))

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handerls.Repo.Home)
	mux.Get("/about", handerls.Repo.About)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux

}
