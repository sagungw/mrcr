package route

import (
	"sagungw/mercari/api/handler"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

func RouteHandler(httpHandler handler.Handler) *chi.Mux {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"*"},
	})

	r.Use(cors.Handler)
	r.Use(handler.RequestLogger())
	r.Get("/api/provinces", httpHandler.GetProvinces)
	r.Get("/api/cities", httpHandler.GetCities)
	r.Get("/api/districts", httpHandler.GetDistricts)
	r.Get("/api/subdistricts", httpHandler.GetSubDistricts)

	return r
}
