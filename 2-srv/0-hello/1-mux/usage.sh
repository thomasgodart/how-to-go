#!/bin/bash

go run . &

sleep 2

go test . &&
{
wget -q -O - http://how-srv-hello:8080/      # Should display: "Hello "
wget -q -O - http://how-srv-hello:8080/World # Should display: "Hello World"
}
pkill 1-mux
