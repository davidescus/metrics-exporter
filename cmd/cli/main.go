package main

import (
	"exporter-imporoved/pkg/app"
	"exporter-imporoved/pkg/collectors"
	"exporter-imporoved/pkg/config"
	"exporter-imporoved/pkg/storage"
	"log"
	"time"
)

func main() {
	// create config
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	// create storage
	store, err := storage.NewMemory()
	if err != nil {
		log.Fatal(err)
	}

	// create app
	application := app.NewApp(conf, store)

	// register collectors
	timeout := time.Duration(conf.MaxProcessCollectDuration) * time.Second
	// TODO read all partitions and assign a collector for each one
	application.RegisterCollector(collectors.NewDisk("sda1", timeout))
	application.RegisterCollector(collectors.NewRam(timeout))

	// listen signals for graceful shutdown
	go application.ListenSignals()

	// run application
	application.Run()
}
