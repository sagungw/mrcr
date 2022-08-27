package route

import (
	"sagungw/mercari/api/handler"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

func RouteHandler(httpHandler handler.Handler, httpMiddleware handler.Middleware) *chi.Mux {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"*"},
	})

	r.Use(cors.Handler)
	r.Use(httpMiddleware.RequestLogger())
	r.Put("/api/users/register", httpHandler.Register)
	r.Post("/api/users/login", httpHandler.Login)
	r.With(httpMiddleware.Auth()).Get("/api/users/login/history", httpHandler.LoginHistory)

	return r
}
