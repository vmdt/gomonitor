package main

import (
	"consumer/config"
	"consumer/socket"
	"fmt"
	"net/http"
)

func main() {
	var cf = config.NewConfig()
	fmt.Printf("Setting up server in: %s", cf.CONSUMER_PORT)
	hub := socket.NewHub()
	sws := socket.NewServerWS(hub)
	go hub.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		sws.StartServerWS(hub, w, r)
	})

	http.ListenAndServe(fmt.Sprintf(":%s", cf.CONSUMER_PORT), nil)
}
