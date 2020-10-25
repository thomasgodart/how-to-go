package main

import (
	"log"
	"net/http"
	"time"

	"github.com/coreos/go-systemd/activation"
	"github.com/coreos/go-systemd/daemon"
)

// from https://vincent.bernat.ch/en/blog/2018-systemd-golang-socket-activation

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

	listeners, err := activation.Listeners()
	if err != nil {
		log.Panicf("activation.Listeners(true) error: %s", err)
	}
	if len(listeners) != 1 {
		log.Panicf("unexpected number of socket activation (%d != 1)", len(listeners))
	}

	notify() // notify must always be placed after "listen" and before "serve"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("system" + "\n"))
	})

	http.Serve(listeners[0], nil)
}
