package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Websocket struct{}

func NewWebsocket() *Websocket {
	return &Websocket{}
}

func (wh *Websocket) Handle(w http.ResponseWriter, r *http.Request) {
	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	if _, err := upgrader.Upgrade(w, r, nil); err != nil {
		log.Fatal(err)
	}
}
