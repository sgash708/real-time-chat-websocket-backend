package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sgash708/real-time-chat-websocket-backend/api/domain/model"
)

type Websocket struct {
	Hub *model.Hub
}

func NewWebsocket(hub *model.Hub) *Websocket {
	return &Websocket{
		Hub: hub,
	}
}

func (wh *Websocket) Handle(w http.ResponseWriter, r *http.Request) {
	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := model.NewClient(ws)
	go client.ReadLoop(wh.Hub.BroadcastCh, wh.Hub.UnRegisterCh)
	go client.WriteLoop()
	wh.Hub.RegisterCh <- client
}
