package middlewares

import (
	"compress/gzip"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// UseGZip uses middleware for compressing response using gzip algorithm.
func UseGZip() Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			encodings := r.Header.Get("Accept-Encoding")
			if !strings.Contains(encodings, "gzip") {
				h.ServeHTTP(w, r)
				return
			}
			w.Header().Add("Content-Encoding", "gzip")
			gw := gzip.NewWriter(w)
			defer gw.Close()
			grw := gzipResponseWriter{
				ResponseWriter: w,
				Writer:         gw,
			}
			h.ServeHTTP(&grw, r)
		})
	}
}

type gzipResponseWriter struct {
	http.ResponseWriter
	io.Writer
	status int
}

func (grw *gzipResponseWriter) WriteHeader(statusCode int) {
	grw.status = statusCode
	if grw.status == 0 {
		grw.status = 200
	}
	grw.ResponseWriter.WriteHeader(statusCode)
}

func (grw *gzipResponseWriter) Write(data []byte) (int, error) {
	if len(data) > 1400 || grw.status >= 400 {
		return grw.Writer.Write(data)
	}
	grw.Header().Del("Content-Encoding")
	grw.Header().Add("Content-Length", strconv.Itoa(len(data)))
	return grw.ResponseWriter.Write(data)
}
