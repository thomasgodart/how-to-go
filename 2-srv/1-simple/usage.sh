#!/bin/bash

go run *.go &

sleep 2

wget -q -O - http://how-srv-simple:8080/
wget -q -O - http://how-srv-simple:8080/aaa
wget -q -O - http://how-srv-simple:8080/bbb

pkill main
