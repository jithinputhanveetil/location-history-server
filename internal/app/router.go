package app

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// InitRouter sets up location-history router.
func (s *Server) InitRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Mount("/location", s.initLocationHistoryRouter())
	r.NotFound(s.handle404)
	return r
}

// handle404 sets up custom 404 handling.
func (s *Server) handle404(w http.ResponseWriter, r *http.Request) {
	fail(
		w,
		http.StatusNotFound,
		"OOPS. We tried 404 times, but couldn't find that resource.",
	)
}
