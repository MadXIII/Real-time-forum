package server

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

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
		Clients:    make(map[*Client]bool),
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
	fmt.Println("W", w)
	fmt.Println("R", r)
	conn, err := upgrader.Upgrade(w, r, nil)
	log.Println(conn)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: h, conn: conn, send: make(chan []byte, 256)}
	client.hub.Registr <- client
	go client.WriteClient()
	go client.ReadClient()
}

func (c *Client) WriteClient() {
	ticker := time.NewTicker(55 * time.Second)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Println(err)
				return
			}

			w.Write(message)

			for i := 0; i < len(c.send); i++ {
				w.Write([]byte{'\n'})
				w.Write(<-c.send)
			}

			if err = w.Close(); err != nil {
				log.Println(err)
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func (c *Client) ReadClient() {
	defer func() {
		c.hub.Unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(512)
	c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Websocket err: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, []byte{'\n'}, []byte{' '}, -1))
		c.hub.Broadcast <- message
	}
}
