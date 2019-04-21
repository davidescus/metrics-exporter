package collectors

import (
	"errors"
	"exporter-imporoved/pkg/metric"
	"time"
)

// Ram ...
type Disk struct {
	name    string
	mount   string
	timeout time.Duration
}

func NewDisk(mount string, timeout time.Duration) *Disk {
	return &Disk{
		name:    "Disk",
		mount:   mount,
		timeout: timeout,
	}
}

// Collect ...
func (d *Disk) Collect() (*metric.Metric, error) {
	c1 := make(chan *metric.Metric, 1)
	go func() {
		// TODO deal with panic here
		// TODO write implementation
		// fake metric
		m := metric.NewMetric(d.name+"-"+d.mount, "234234234", time.Now())
		c1 <- m
	}()

	select {
	case m := <-c1:
		return m, nil
	case <-time.After(d.timeout):
		return nil, errors.New("failed to read" + d.name + "-" + d.mount + " metric")
	}
}
