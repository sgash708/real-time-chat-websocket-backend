package handler

import (
	"net/http"

	"github.com/sgash708/real-time-chat-websocket-backend/api/client/websocket"
)

type Handler struct {
	WebSocket *websocket.Websocket
}

func NewHandler(
	ws *websocket.Websocket,
) *Handler {
	return &Handler{
		WebSocket: ws,
	}
}

func (h *Handler) Routes() {
	http.HandleFunc("/ws", h.WebSocket.Handle)
}
