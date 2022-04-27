package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Hub struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Registr    chan *Client
	Unregister chan *Client
}

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

func NewHub() *Hub {
	return &Hub{
		Clients:    map[*Client]bool{},
		Broadcast:  make(chan []byte),
		Registr:    make(chan *Client),
		Unregister: make(chan *Client),
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WSChat(h *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	client := &Client{hub: h, conn: conn, send: make(chan []byte, 256)}
	client.hub.Registr <- client
	go client.WriteClient()
	go client.ReadClient()
}

func (c *Client) WriteClient() {
}

func (c *Client) ReadClient() {
}
