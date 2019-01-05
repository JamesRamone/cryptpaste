package main

func (s *server) routes() {
	s.router.Get("/api/pad/{id}", s.handleGetPad())
	s.router.Post("/api/pad", s.handleCreatePad())
	s.router.Get("/", s.handleIndex())
}
