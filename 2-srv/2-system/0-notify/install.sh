#!/bin/bash

go get github.com/coreos/go-systemd/daemon

cd /etc/systemd/system/

sudo ln -sf /var/www/go/how-to-go/2-srv/2-system/0-notify/system.service how-srv-system0.service

sudo systemctl daemon-reload

sudo systemctl start how-srv-system0.service

sudo systemctl enable how-srv-system0.service
