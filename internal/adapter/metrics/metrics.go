package metrics

import (
    "github.com/prometheus/client_golang/prometheus"
)

type PrometheusMetrics struct {
    httpRequests *prometheus.CounterVec
}

func NewPrometheusMetrics() *PrometheusMetrics {
    m := &PrometheusMetrics{
        httpRequests: prometheus.NewCounterVec(
            prometheus.CounterOpts{
                Name: "http_requests_total",
                Help: "Total HTTP requests by status, method, and path",
            },
            []string{"method", "path", "status"},
        ),
    }

    prometheus.MustRegister(m.httpRequests)
    return m
}

func (p *PrometheusMetrics) IncHttpRequest(method, path, status string) {
    p.httpRequests.WithLabelValues(method, path, status).Inc()
}
