package collectors

import "exporter-imporoved/pkg/metric"

// Collecter must be implement by all collectors
type Collecter interface {
	Collect() (*metric.Metric, error)
}
