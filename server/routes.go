package main

import "github.com/go-chi/cors"

func (s *server) routes() {
	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	s.router.Use(cors.Handler)

	s.router.Get("/api/pad/{id}", s.handleGetPad())
	s.router.Post("/api/pad", s.handleCreatePad())
	//s.router.Get("/*", s.handleIndex())
}
