package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type room struct {
	//他のクライアントに転送するためのメッセージを保持するチャネル
	forword chan []byte
	// チャットに参加するクライアントのためのチャネル
	join chan *client
	// チャットから退室するクライアントのためのチャネル
	leave chan *client
	// 在室している全てのクライアントのリスト
	clients map[*client]bool
}

func (r *room) run() {
	for {
		select {
		case client := <-r.join:
			fmt.Println("join")
			r.clients[client] = true
		case client := <-r.leave:
			fmt.Println("leave")
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forword:
			fmt.Println("announce")
			announceMessageForAll(r, msg)
		}
	}
}

func announceMessageForAll(r *room, msg []byte) {
	for client := range r.clients {
		select {
		case client.send <- msg:
			fmt.Println("send")
			//メッセージを送信
		default:
			fmt.Println("close")
			delete(r.clients, client)
			close(client.send)
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}
	r.join <- client

	leaveFunc := func() { r.leave <- client }
	defer leaveFunc()

	go client.write()
	client.read()

}

func newRoom() *room {
	return &room{
		forword: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}
