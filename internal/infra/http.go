package infra

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"time"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Compress(6))

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))
	r.Use(middleware.Heartbeat("/health"))
	r.Use(middleware.Compress(6))

	return r
}