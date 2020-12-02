package middleware

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/http"
)

var (
	HttpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "myapp_http_duration_seconds",
		Help: "Duration of HTTP requests.",
	}, []string{"path"})
	HttpHits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "myapp_http_hits",
		Help: "Hits of HTTP requests.",
	}, []string{"path"})
)

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()
		timer := prometheus.NewTimer(HttpDuration.WithLabelValues(path))
		HttpHits.WithLabelValues(path).Inc()
		next.ServeHTTP(w, r)
		timer.ObserveDuration()
	})
}
