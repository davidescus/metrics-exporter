package storage

import "exporter-imporoved/pkg/metric"

type Storable interface {
	Store(metric metric.Metric) error
	GetAll() ([]metric.Metric, error)
}
