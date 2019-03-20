#!/usr/bin/env bash

# create user
sudo useradd --no-create-home --shell /bin/false metrics-exporter

# create directories
sudo mkdir /etc/metrics-exporter
sudo mkdir /var/lib/metrics-exporter

# change owner and group
sudo chown metrics-exporter:metrics-exporter /etc/metrics-exporter
sudo chown metrics-exporter:metrics-exporter /var/lib/metrics-exporter

# move files on their place (from /tmp)
sudo mv metrics-exporter /usr/local/bin/
sudo mv metrics-exporter.yaml /etc/metrics-exporter/
sudo mv metrics-exporter.service /etc/systemd/system/metrics-exporter.service

# remove installer
rm -rf metrics-exporter.sh

# change owner and group for moved files
sudo chown metrics-exporter:metrics-exporter /usr/local/bin/metrics-exporter
sudo chown metrics-exporter:metrics-exporter /etc/metrics-exporter/metrics-exporter.yaml

# restart daemon and start service
sudo systemctl daemon-reload
sudo systemctl start metrics-exporter