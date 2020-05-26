package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type Broadcast struct {
	source  *Client
	message []byte
}

type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan *Broadcast
	register   chan *Client
	unregister chan *Client
}
type Client struct {
	socket net.Conn
	data   chan []byte
}
func fmtDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("%02d:%02d", h, m)
}

func (manager *ClientManager) start() {
	for {
		select {
		case connection := <-manager.register:
			manager.clients[connection] = true
			log.Println("Added new connection!")
		case connection := <-manager.unregister:
			if _, ok := manager.clients[connection]; ok {
				close(connection.data)
				delete(manager.clients, connection)
				log.Println("A connection has terminated!")
			}
		case broadcast := <-manager.broadcast:
			message := broadcast.message
			for connection := range manager.clients {
				if connection == broadcast.source {
					continue
				}
				select {
				case connection.data <- message:
				default:
					close(connection.data)
					delete(manager.clients, connection)
				}
			}
		}
	}
}

func (manager *ClientManager) receive(client *Client) {
	for {
		message := make([]byte, 4096)
		length, err := client.socket.Read(message)
		if err != nil {
			manager.unregister <- client
			client.socket.Close()
			break
		}
		if length > 0 {
			message = message[0:length]
			log.Println("RECEIVED: " + strings.Trim(string(message), "\r\n"))
			manager.analyze(client, message)
		}
	}
}

func (manager *ClientManager)analyze(source *Client, message []byte) {
	listener := r2queryListener {}

	modTime := time.Now()
	listener.execute(string(message))
	since := time.Since(modTime)
	listener.result += fmt.Sprintln("{",since,"}")
	source.data <- []byte(listener.result)

	//str := fmt.Sprintf("%d\n", listener.pop())
	//source.data <- []byte(str)
	//manager.broadcast <- &Broadcast{source: source, message: message}
}
func (manager *ClientManager) send(client *Client) {
	defer client.socket.Close()
	for {
		select {
		case message, ok := <-client.data:
			if !ok {
				return
			}
			client.socket.Write(message)
		}
	}
}
func main() {
	address := os.Getenv("R2_ADDRESS")
	if address == "" {
		address = ":12345"
	}
	log.Printf("Starting server ... '%s'\n", address)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Println(err)
	}
	manager := ClientManager{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan *Broadcast),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
	go manager.start()
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		client := &Client{socket: connection, data: make(chan []byte)}
		manager.register <- client
		go manager.receive(client)
		go manager.send(client)
	}

}
