package metric

import (
	"fmt"
	"github.com/bongerka/metricaler/internal/metric"
	def "github.com/bongerka/metricaler/internal/repository"
	"gitlab.com/bongerka/lg"
)

var _ def.MetricRepository = (*repository)(nil)

type repository struct {
	conn map[string]*metric.Metric
}

func NewRepository(size int) *repository {
	return &repository{
		conn: make(map[string]*metric.Metric, size),
	}
}

func (r *repository) ChangeCounter(name string, value int64) {
	newName := fmt.Sprintf("%s_%s", metric.Counter, name)

	prev, ok := r.conn[newName]
	if !ok {
		prev = &metric.Metric{Type: metric.Counter, Value: int64(0)}
	}
	newValue, ok := prev.Value.(int64)
	if !ok {
		lg.Errorf("counter is not int64 but %T", prev.Value)
		return
	}

	r.conn[newName] = &metric.Metric{
		Type:  metric.Counter,
		Value: newValue + value,
	}
}

func (r *repository) ChangeGauge(name string, value float64) {
	newName := fmt.Sprintf("%s_%s", metric.Gauge, name)
	r.conn[newName] = &metric.Metric{
		Type:  metric.Gauge,
		Value: value,
	}
}
