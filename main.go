package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"metrics-exporter/conf"
	"metrics-exporter/metric"
	"metrics-exporter/metric-type/disk"
	"metrics-exporter/persistence"
	"net/http"
	"strconv"
	"time"
)

var configPath = "/etc/metrics-exporter/metrics-exporter.yaml"

func main() {
	// Rewrite config path using startup args
	c := conf.NewConf(configPath)

	fmt.Printf("Webserver start at: %v\n", c.Port)
	db.CreateSchemaIfNotExists(c)

	// start collect metrics in separate goroutine
	go collectMetrics(c)

	// execute cleanup in separate goroutine
	go cleanExpiredMetrics(c)

	handleRoute()
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(c.Port), nil))
}

func collectMetrics(conf *conf.Conf) {
	for true {
		// run all metrics
		m := disk.Collect()
		db.StoreMetric(m, conf)
        // another metric ...

		time.Sleep(time.Second * time.Duration(conf.CollectInterval))
	}
}

func cleanExpiredMetrics(conf *conf.Conf) {
	for true {
		db.RemoveRecordsOlderThan(conf)
		time.Sleep(time.Second * time.Duration(conf.TimeToLive))
	}
}

func handleRoute() {
	http.HandleFunc("/metrics", get)
}

type Response struct {
	Fqdn string `json:"fqdn"`
	Metrics metric.Metrics `json:"metrics"`
}

func get(w http.ResponseWriter, r *http.Request) {

	paramUnixTimestamp := r.URL.Query().Get("fromUnixTimestamp")
	// assume it wants all records if param not exists
	if paramUnixTimestamp == "" {
		paramUnixTimestamp = "0"
	}

	unixTimestamp, err := strconv.ParseUint(paramUnixTimestamp, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	c := conf.NewConf(configPath)
	response := Response{}
	response.Fqdn = "TODO here must be unique identifier of system"
	response.Metrics = db.GetFromDate(c, unixTimestamp)

	jsonData, encodeError := json.Marshal(response)
	if encodeError != nil {
		log.Fatal(encodeError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(jsonData)
}



