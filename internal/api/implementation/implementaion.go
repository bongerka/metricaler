package implementation

import "github.com/bongerka/metricaler/internal/api/implementation/metric"

type Implemetation struct {
	MetricService *metric.Service
}

func NewImplementation(ms *metric.Service) *Implemetation {
	return &Implemetation{
		MetricService: ms,
	}
}
