package disk

import (
	"github.com/shirou/gopsutil/disk"
	"log"
	"metrics-exporter/metric"
	"strconv"
	"time"
)

const metricType = "Physical Disk Usage"
const metricName = "main_disk"

func Collect() *metric.Metric {

	m := &metric.Metric{}

	diskStat, err := disk.Usage("/")
	if err != nil {
		log.Fatal(err)
	}

	m.Type = metricType
	m.Name = metricName
	m.Total = strconv.FormatUint(diskStat.Total, 10)
	m.InUse = strconv.FormatUint(diskStat.Used, 10)
	m.UnixTimestamp = uint64(time.Now().Unix())

	return m
}