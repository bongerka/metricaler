package api

import (
	"github.com/bongerka/metricaler/internal/api/implementation/metric"
	"net/http"
)

func MapHandlers(mux *http.ServeMux, service *metric.Service) error {
	mux.HandleFunc("/", service.PostMetric)

	return nil
}
