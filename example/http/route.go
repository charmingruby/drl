package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(
	r *chi.Mux,
	mw Middleware,
) {

	r.Get("/", mw.rateLimiter(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	}))
}
