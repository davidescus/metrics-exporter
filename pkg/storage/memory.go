package storage

import "exporter-imporoved/pkg/metric"

type Memory struct {
	data []metric.Metric
}

func NewMemory() (*Memory, error) {
	memory := Memory{
		data: make([]metric.Metric, 0, 0),
	}

	return &memory, nil
}

// Store ...
func (m *Memory) Store(metric metric.Metric) error {
	m.data = append(m.data, metric)
	return nil
}

// GetAll
func (m *Memory) GetAll() ([]metric.Metric, error) {
	return m.data, nil
}
