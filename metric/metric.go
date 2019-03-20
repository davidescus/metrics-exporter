package metric

type Metrics []Metric

type Metric struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Total string `json:"total"`
	InUse string `json:"inUse"`
	UnixTimestamp  uint64 `json:"unixTimestamp"`
}
