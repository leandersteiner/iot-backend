package server

import (
	"github.com/go-chi/chi/v5"
)

func newRouter() chi.Router {
	r := chi.NewRouter()

	configureMiddleware(r)

	r.Route("/api", func(r chi.Router) {
		r.Get("/", healthCheck)
		r.Get("/heartrate", getAllHeartRates)
		r.Get("/heartrate/{count}", getHeartRateRange)
		r.Post("/heartrate", postHeartRate)
	})

	return r
}
