package chat

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
)

type Client struct {
	ClientName string
	Ws         *websocket.Conn
	//Send       chan string
}

func NewClient(ws *websocket.Conn, name string) *Client {
	//	send := make(chan string)
	return &Client{name, ws}
}

func (c *Client) PushMessage(message string) {
	//b := <-c.Send
	if err := websocket.Message.Send(c.Ws, message); err != nil {
		fmt.Println("Can't send")
		return
	}
}
