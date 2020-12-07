#!/bin/bash

go run . &

sleep 2

go test . &&
{
wget -q -O - http://how-srv-simple:8080/
wget -q -O - http://how-srv-simple:8080/aaa
wget -q -O - http://how-srv-simple:8080/bbb
}
pkill 1-simple
