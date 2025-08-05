package http

import (
	"log/slog"
	"net"
	"net/http"

	"github.com/charmingruby/drl/pkg/rate_limiter"
)

type Middleware struct {
	rl  *rate_limiter.RateLimiter
	log *slog.Logger
}

func NewMiddleware(rl *rate_limiter.RateLimiter, log *slog.Logger,
) Middleware {
	return Middleware{
		log: log,
		rl:  rl,
	}
}

func (mw *Middleware) rateLimiter(
	next http.HandlerFunc,
) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP, _, _ := net.SplitHostPort(r.RemoteAddr)

		allowed, err := mw.rl.Allow(clientIP)

		if err != nil {
			mw.log.Error("unable to validate rate limiting", "error", err)

			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		if !allowed {
			mw.log.Debug("unable to pass by the rate limiting", "ip", clientIP)

			http.Error(w, "too many requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
