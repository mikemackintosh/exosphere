package exo

import (
	"net/http"
	"time"

	stdslog "golang.org/x/exp/slog"
)

type statusRecorder struct {
	http.ResponseWriter
	startTime time.Time
	Status    int
	Size      int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.Status = code
	rec.ResponseWriter.WriteHeader(code)
}

func (rec *statusRecorder) Write(w []byte) (int, error) {
	rec.Size = len(w)
	return rec.ResponseWriter.Write(w)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Initialize the status to 200 in case WriteHeader is not called
		rec := statusRecorder{ResponseWriter: w, Status: 200, Size: 0, startTime: time.Now()}
		next.ServeHTTP(&rec, r)

		slog.WithGroup("http").LogAttrs(
			r.Context(),
			stdslog.LevelInfo,
			"request served",
			[]stdslog.Attr{
				stdslog.Any("duration", (time.Since(rec.startTime).String())),
				stdslog.Any("duration_milli", time.Since(rec.startTime).Milliseconds()),
				stdslog.Any("remote_addr", r.RemoteAddr),
				stdslog.Any("method", r.Method),
				stdslog.Any("uri", r.RequestURI),
				stdslog.Any("proto", r.Proto),
				stdslog.Any("encoding", r.TransferEncoding),
				stdslog.Any("response_code", rec.Status),
				stdslog.Any("response_size", rec.Size),
				stdslog.Any("Content-length", r.ContentLength),
				stdslog.Any("request-headers", r.Header),
				stdslog.Any("headers", r.Header),
				stdslog.Any("Content-type", r.Header.Get("Content-type")),
				stdslog.String("user_agent", r.UserAgent()),
			}...,
		)
	})
}
