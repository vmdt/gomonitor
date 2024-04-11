package socket

import (
	"consumer/config"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	cf            = config.NewConfig()
	kafkaConsumer = config.NewConsumer()
)

type ServerWS struct {
	Hub      *Hub
	Upgrader websocket.Upgrader
}

func NewServerWS(hub *Hub) *ServerWS {
	return &ServerWS{
		Hub: hub,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // accept all requests
			},
		},
	}
}

func (sws *ServerWS) StartServerWS(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := sws.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	go client.WritePump()
	go kafkaConsumer.ReadMessage(cf.KAFKA_TOPIC, client.WatchMessage)
}
