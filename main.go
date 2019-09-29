package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"net/http"
	"os"
	"rayyildiz.dev/app/internal/infra"
	"time"
)

func init() {
	godotenv.Load()
}

func main() {
	log := infra.NewLogger()
	defer sentry.Flush(time.Second * 5)
	db, err := infra.NewDatabase()
	if err != nil {
		log.Fatal("could not initialize database", zap.Error(err))
	}
	defer db.Close()

	r := infra.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	// Handlers
	r.Route("/api", func(r chi.Router) {

	})

	log.Info("server is starting", zap.String("port", port))
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Error("could not start server", zap.String("port", port), zap.Error(err))
	}
}