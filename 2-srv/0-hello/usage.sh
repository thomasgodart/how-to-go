#!/bin/bash

go run *.go

wget -q -O - http://how-srv-hello:8080/      # Should display: "Hello "
wget -q -O - http://how-srv-hello:8080/World # Should display: "Hello World"
