package model

import (
	"log"

	"github.com/gorilla/websocket"
)

// Client websocketを使ってメッセージをやりとりする
type Client struct {
	ws     *websocket.Conn
	sendCh chan []byte // メッセージのやり取りに使う
}

func NewClient(ws *websocket.Conn) *Client {
	return &Client{
		ws:     ws,
		sendCh: make(chan []byte),
	}
}

func (c *Client) ReadLoop(broadCast chan<- []byte, unregister chan<- *Client) {
	defer func() {
		c.disconnect(unregister)
	}()

	for {
		_, jsonMsg, err := c.ws.ReadMessage()
		if err != nil && websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			log.Println(err)
			break
		}

		// 読み取り成功後、broadCastに読み取った値を送信する
		broadCast <- jsonMsg
	}
}

func (c *Client) WriteLoop() {
	defer func() {
		// websocketコネクションがあるブラウザにレスポンスを返す
		c.ws.Close()
	}()

	for {
		// *Clientが持つsendChに値が送信されるのを待つ
		message := <-c.sendCh

		// websocketレスポンスをするためのWriterを発行
		w, err := c.ws.NextWriter(websocket.TextMessage)
		if err != nil {
			log.Println(err)
			return
		}
		// Writeにmessageを書き込む
		w.Write(message)

		if err := w.Close(); err != nil {
			log.Println(err)
			return
		}
	}
}

func (c *Client) disconnect(unregister chan<- *Client) {
	unregister <- c
	c.ws.Close()
}
