package model

// Hub ユーザの入退出やチャット送信を管理する
type Hub struct {
	Clients      map[*Client]bool // 現在のチャット参加者一覧
	RegisterCh   chan *Client     // Clientを参照。ユーザ入室時に利用
	UnRegisterCh chan *Client     // Clientを参照。ユーザ退室時に利用
	BroadcastCh  chan []byte      // []byteのやりとりをする。ユーザがチャット送信時に利用
}

func NewHub() *Hub {
	return &Hub{
		Clients:      make(map[*Client]bool),
		RegisterCh:   make(chan *Client),
		UnRegisterCh: make(chan *Client),
		BroadcastCh:  make(chan []byte),
	}
}

func (h *Hub) RunLoop() {
	for {
		select {
		case client := <-h.RegisterCh:
			h.register(client)
		case client := <-h.UnRegisterCh:
			h.unregister(client)
		case msg := <-h.BroadcastCh:
			h.broadCastToAllClient(msg)
		}
	}
}

// register 入室
func (h *Hub) register(c *Client) {
	h.Clients[c] = true
}

// register 退室
func (h *Hub) unregister(c *Client) {
	delete(h.Clients, c)
}

// broadCastToAllClient レシーバ(Hub)の全ClientのsendChへ値を送信
func (h *Hub) broadCastToAllClient(msg []byte) {
	for c := range h.Clients {
		c.sendCh <- msg
	}
}
