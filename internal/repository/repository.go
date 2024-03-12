package repository

type MetricRepository interface {
	ChangeCounter(name string, value int64)
	ChangeGauge(name string, value float64)
}
