package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/coreos/go-systemd/daemon"
)

// from https://vincent.bernat.ch/en/blog/2017-systemd-golang

func notify() {

	daemon.SdNotify(false, daemon.SdNotifyReady)
	go func() {
		interval, err := daemon.SdWatchdogEnabled(false)
		if err != nil || interval == 0 {
			return
		}
		for {
			daemon.SdNotify(false, daemon.SdNotifyWatchdog)
			time.Sleep(interval / 3)
		}
	}()
}

func main() {

	listen, err := net.Listen("tcp", "how-srv-system0:8080")
	if err != nil {
		log.Panicf("net.Listen(\"tcp\", \"how-srv-system0:8080\") error: %s", err)
	}

	notify() // notify must always be placed after "listen" and before "serve"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("system" + "\n"))
	})

	http.Serve(listen, nil)
}
