package collectors

import (
	"errors"
	"exporter-imporoved/pkg/metric"
	"time"
)

// Ram ...
type Ram struct {
	name    string
	timeout time.Duration
}

// NewRam ...
func NewRam(timeout time.Duration) *Ram {
	return &Ram{
		name:    "Ram",
		timeout: timeout,
	}
}

// Collect ...
func (r *Ram) Collect() (*metric.Metric, error) {
	c1 := make(chan *metric.Metric, 1)
	go func() {
		// TODO deal with panic here
		// TODO write implementation
		// fake metric
		m := metric.NewMetric("", "234234234", time.Now())
		c1 <- m
	}()

	select {
	case m := <-c1:
		return m, nil
	case <-time.After(r.timeout):
		return nil, errors.New("failed to read" + r.name + " metric")
	}
}
