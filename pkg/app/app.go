package app

import (
	"exporter-imporoved/pkg/collectors"
	"exporter-imporoved/pkg/config"
	"exporter-imporoved/pkg/server"
	"exporter-imporoved/pkg/storage"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type app struct {
	closeChan  chan bool
	config     *config.Config
	storage    storage.Storable
	collectors []collectors.Collecter
}

// NewApp will create application
func NewApp(conf *config.Config, store storage.Storable) *app {
	return &app{
		closeChan:  make(chan bool),
		config:     conf,
		storage:    store,
		collectors: make([]collectors.Collecter, 0, 0),
	}
}

// RegisterCollector will append new collector to app
func (a *app) RegisterCollector(c collectors.Collecter) {
	a.collectors = append(a.collectors, c)
}

// Run will start application
func (a *app) Run() {
	// start webserver
	// TODO graceful shutdown webserver
	s := server.NewServer(a.config.Port)
	go s.Start(a.storage)

	// run collectors or shutdown
	for {
		select {
		case <-a.closeChan:
			log.Println("Shutdown. Have a good day!")
			return
		default:
			a.runCollectors()
		}

		// TODO use time.Timer (ask Adrian)
		time.Sleep(time.Duration(a.config.ScrapeInterval) * time.Second)
	}
}

// runCollectors will trigger all collectors
func (a *app) runCollectors() {
	// TODO implement goroutine for each collector
	for _, v := range a.collectors {
		m, err := v.Collect()
		if err != nil {
			// TODO log this, NO fatal
			log.Println(err)
			continue
		}

		err = a.storage.Store(*m)
		if err != nil {
			// TODO log this NO fatal
			log.Println(err)
			continue
		}
	}
}

// ListenSignals will take care about all signals,
// it is used for graceful shutdown
func (a *app) ListenSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-c:
		a.Close()
	}
}

// Close will close all channels
func (a *app) Close() {
	close(a.closeChan)
}
