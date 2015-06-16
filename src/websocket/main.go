package main

import (
	. "chat"
	"code.google.com/p/go.net/websocket"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func chat(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	t, _ := template.ParseFiles("client.html")
	t.Execute(w, nil)
}
func join(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	t, _ := template.ParseFiles("join.html")
	t.Execute(w, nil)
}
func main() {
	server := NewServer()
	go server.Start()
	http.Handle("/ws", websocket.Handler(server.OnConnected))
	http.HandleFunc("/chat", chat)
	http.HandleFunc("/", join)
	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal("ListentAndServe:", err)
	}
}
