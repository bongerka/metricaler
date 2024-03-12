package metric

type MetType string

const (
	Gauge   MetType = "gauge"
	Counter MetType = "counter"
)

type Metric struct {
	Type  MetType
	Value any
}
