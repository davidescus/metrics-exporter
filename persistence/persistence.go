package db

import (
	"database/sql"
	"log"
	"metrics-exporter/conf"
	"metrics-exporter/metric"
	"time"
)

// connect to storage
func connect(storagePath string) *sql.DB {
	resource, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		log.Fatal(err)
	}
	return resource
}

// create schema and table if not exists
func CreateSchemaIfNotExists(conf *conf.Conf) {
	resource := connect(conf.StoragePath)
    defer resource.Close()

	query := `
        CREATE TABLE IF NOT EXISTS metrics (
            Id INTEGER PRIMARY KEY,
            Type TEXT,
            Name TEXT,
            Total TEXT, 
            InUse TEXT,
            UnixTimestamp TEXT
        )`
	stmt, err := resource.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec()
}

// store metric
func StoreMetric(m *metric.Metric, conf *conf.Conf) {
	resource := connect(conf.StoragePath)
    defer resource.Close()

	query := "INSERT INTO metrics (Type, Name, Total, InUse, UnixTimestamp) VALUES (?, ?, ?, ?, ?)"
	stmt, err := resource.Prepare(query)
	if err != nil {
		log.Println(err)
	}

	_, err = stmt.Exec(m.Type, m.Name, m.Total, m.InUse, m.UnixTimestamp)
	if err != nil {
		log.Println(err)
	}
}

// remove records that exceed time to live
func RemoveRecordsOlderThan(conf *conf.Conf) {
	resource := connect(conf.StoragePath)
	defer resource.Close()

	query := "DELETE FROM metrics WHERE UnixTimestamp < ?"

	stmt, err := resource.Prepare(query)
	if err != nil {
		log.Println(err)
	}

	maxRecordTimeToLive := uint64(time.Now().Unix()) - conf.TimeToLive
	_, err = stmt.Exec(maxRecordTimeToLive)
	if err != nil {
		log.Println(err)
	}
}

func GetFromDate(conf *conf.Conf, fromDate uint64) metric.Metrics {
	storage := connect(conf.StoragePath)
	defer storage.Close()

	query := "SELECT Type, Name, Total, InUse, UnixTimestamp FROM metrics WHERE UnixTimestamp > ?"
	rows, err := storage.Query(query, fromDate)
	if err != nil {
		log.Fatal("Something went wrong when query database for records ...")
	}

	metrics := metric.Metrics{}
	for rows.Next() {
		m := &metric.Metric{}
		err := rows.Scan(&m.Type, &m.Name, &m.Total, &m.InUse, &m.UnixTimestamp)
		if err != nil {
			log.Fatal(err)
		}
		metrics = append(metrics, *m)
	}

	return metrics
}