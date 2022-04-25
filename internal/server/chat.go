package server

import "github.com/gorilla/websocket"

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
