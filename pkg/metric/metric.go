package metric

import "time"

type Metric struct {
	Name      string    `json:"Name"`
	Value     string    `json:"Value"`
	Timestamp time.Time `json:"Timestamp"`
}

// NewMetric return pointer to Metric instance
func NewMetric(name string, value string, timestamp time.Time) *Metric {
	return &Metric{
		Name:      name,
		Value:     value,
		Timestamp: timestamp,
	}
}
