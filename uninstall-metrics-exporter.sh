#!/usr/bin/env bash

# stop, disable service
sudo systemctl stop metrics-exporter
sudo systemctl disable metrics-exporter

# delete user
sudo userdel -r metrics-exporter

# remove all files
rm -rf metrics-exporter*
rm -rf uninstall-metrics-exporter.sh

sudo rm -rf /etc/metrics-exporter
sudo rm -rf /var/lib/metrics-exporter

# remove service
sudo rm -rf /etc/systemd/system/metrics-exporter.service

# restart daemon and start service
sudo systemctl daemon-reload