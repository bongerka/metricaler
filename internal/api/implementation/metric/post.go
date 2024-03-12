package metric

import (
	"github.com/bongerka/metricaler/internal/metric"
	"gitlab.com/bongerka/lg"
	"net/http"
	"strconv"
	"strings"
)

func (service *Service) PostMetric(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { // TODO where validation should be?
		http.Error(w, "must be POST", http.StatusMethodNotAllowed)
		lg.Warnf("validation: method is %s not POST", r.Method)
		return
	}

	//contentType := r.Header.Get("Content-Type")
	//if contentType != "text/plain" {
	//	http.Error(w, "wrong input format(content-type)", http.StatusBadRequest)
	//	lg.Warnf("validation: %s", contentType)
	//	return
	//}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 5 {
		http.Error(w, "wrong input format(len)", http.StatusNotFound)
		return
	}

	metricType := parts[2]
	metricName := parts[3]
	metricValue := parts[4]

	if metricName == "" {
		http.Error(w, "metric name is disable", http.StatusNotFound)
		return
	}

	if metricType != string(metric.Gauge) && metricType != string(metric.Counter) {
		http.Error(w, "wrong metric type", http.StatusBadRequest)
		return
	}

	if metricType == string(metric.Counter) {
		value, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			http.Error(w, "wrong format of metric value", http.StatusBadRequest)
			return
		}
		service.repo.ChangeCounter(metricName, value)
	}

	if metricType == string(metric.Gauge) {
		value, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			http.Error(w, "wrong format of metric value", http.StatusBadRequest)
			return
		}
		service.repo.ChangeGauge(metricName, value)
	}

	w.WriteHeader(http.StatusOK)
}
