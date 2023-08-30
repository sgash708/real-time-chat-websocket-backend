package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sgash708/real-time-chat-websocket-backend/api/client/websocket"
	"github.com/sgash708/real-time-chat-websocket-backend/api/domain/model"
	"github.com/sgash708/real-time-chat-websocket-backend/api/handler"
)

func main() {
	di()

	port := fmt.Sprintf(":%v", os.Getenv("PORT"))
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}

func di() {
	// Domain
	hub := model.NewHub()
	go hub.RunLoop()

	// Client
	webSocket := websocket.NewWebsocket(hub)

	// Handler
	h := handler.NewHandler(
		webSocket,
	)
	h.Routes()
}
