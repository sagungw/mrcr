package handler

import (
	"net/http"
	"time"

	chi "github.com/go-chi/chi/middleware"
	"github.com/sagungw/gotrunks/log"
	"github.com/sirupsen/logrus"
)

func RequestLogger() func(next http.Handler) http.Handler {
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
