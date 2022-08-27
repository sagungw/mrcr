package handler

import (
	"context"
	"net/http"
	"sagungw/mercari/core/service"
	"strings"
	"time"

	chi "github.com/go-chi/chi/middleware"
	"github.com/sagungw/gotrunks/log"
	"github.com/sirupsen/logrus"
)

type Middleware struct {
	UserService service.UserService
}

func (m *Middleware) RequestLogger() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ww := chi.NewWrapResponseWriter(w, r.ProtoMajor)

			fields := logrus.Fields{
				"remote_addr": r.RemoteAddr,
				"user_agent":  r.UserAgent(),
				"proto":       r.Proto,
				"req_method":  r.Method,
				"req_path":    r.URL.Path,
			}

			start := time.Now()
			defer func() {
				fields["took"] = time.Since(start).Nanoseconds() / 1000000
				fields["res_status"] = ww.Status()
				fields["res_len"] = ww.BytesWritten()

				log.WithFields(fields).Info()
			}()

			next.ServeHTTP(ww, r)
		})
	}
}

func (m *Middleware) Auth() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bearer := r.Header.Get("authorization")
			if bearer == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			split := strings.Split(bearer, " ")
			if len(split) < 2 {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if strings.ToLower(split[0]) != "bearer" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			user, valid := m.UserService.Authorize(r.Context(), split[1])
			if !valid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(context.WithValue(r.Context(), "userToken", split[1]), "userID", user.ID.Hex())
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
