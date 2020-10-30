#!/bin/bash

go get github.com/coreos/go-systemd/activation
go get github.com/coreos/go-systemd/daemon

cd /etc/systemd/system/

sudo ln -sf /var/www/go/how-to-go/2-srv/2-system/1-socket/system.service how-srv-system1.service
sudo ln -sf /var/www/go/how-to-go/2-srv/2-system/1-socket/system.socket how-srv-system1.socket

sudo systemctl daemon-reload

sudo systemctl start how-srv-system1.service
sudo systemctl start how-srv-system1.socket

sudo systemctl enable how-srv-system1.service
sudo systemctl enable how-srv-system1.socket
