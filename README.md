# Metrics Exporter
Collect, store and expose system metrics

## Description
Collect metrics from system, store them for short term (configurable) and expose them as JSON on API endpoint

## Requirements:
-  Linux (tested on Ubuntu 18.04 Bionic Beaver)
-  root permissions 

## How To:
### Install:
```
cd ~ (go to home dir)  
wget https://github.com/davidescus/metrics-exporter/releases/download/v.0.0.1/metrics-exporter.tar.gz  
tar -xzvf metrics-exporter.tar.gz  
./metrics-exporter.sh  
```

### Uninstall:
```
./uninstall-metrics-exporter.sh (on home dir)
```

## Metric types:
* disk usage

## TODO
* implement restriction for maxConnectionNumber (same time)
* add many metric types
* add timeout`s on function call
* add tests
* simplify install process