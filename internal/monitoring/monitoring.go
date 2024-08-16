package monitoring

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupMonitoring() {
	http.Handle("/metrics", promhttp.Handler())
}
