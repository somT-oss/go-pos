package port

type MetricsPort interface {
    IncHttpRequest(method, path, status string)
}
