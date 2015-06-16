package chat

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
)

type Server struct {
	clients   map[string]*Client
	broadcast chan string
	closed    chan bool
}

func NewServer() *Server {
	clients := make(map[string]*Client)
	broadcast := make(chan string)
	closed := make(chan bool)
	return &Server{clients, broadcast, closed}
}

func (s *Server) Start() {
	for {
		select {
		case b := <-s.broadcast:
			//sent broadcast for all users
			fmt.Println("s.broadcast...")
			for _, onlineUser := range s.clients {
				go onlineUser.PushMessage(b)
			}
		case c := <-s.closed:
			if c == true {
				close(s.broadcast)
				close(s.closed)
				return
			}
		}
	}
}

func (s *Server) AddClient(c *Client) {
	(*s).clients[c.ClientName] = c
}
func (s *Server) AddListen() {

}
func (s *Server) OnConnected(ws *websocket.Conn) {
	fmt.Println("OnConnected...")
	username := ws.Request().URL.Query().Get("username")
	fmt.Println(username)
	if "" == username {
		fmt.Println("username is null")
		return
	}
	client := NewClient(ws, username)
	s.AddClient(client)

	for {
		fmt.Println("start")
		var reply string
		if err := websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("can't receive")
			break
		}
		s.broadcast <- reply
	}
	//client quit
	ws.Close()
	delete(s.clients, username)
	//close(this.)
}
