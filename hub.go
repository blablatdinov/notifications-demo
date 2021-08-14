package notifications

import (
	"errors"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type Hub struct {
	MessageChan chan string
	Connections []Connection
}

type Connection struct {
	Username string
	Conn     websocket.Conn
}

func NewHub(ch chan string) Hub {
	return Hub{
		ch,
		[]Connection{},
	}
}

func (h *Hub) existConnection(username string) bool {
	for _, elem := range h.Connections {
		if elem.Username == username {
			return true
		}
	}
	return false
}

func (h *Hub) ConnectUser(conn websocket.Conn, username string) {
	if h.existConnection(username) {
		return
	}
	h.Connections = append(h.Connections, Connection{
		Conn:     conn,
		Username: username,
	})
	log.Printf("Connect user <%s>\n", username)
}

func (h *Hub) Send(username string, message string) error {
	if !h.existConnection(username) {
		return errors.New("user not connected to service")
	}
	h.MessageChan <- message
	return nil
}

func (h *Hub) DisconnectUser(index int) {
	h.Connections = append(h.Connections[:index], h.Connections[index+1:]...)
}

func (h *Hub) PingUsers() {
	go func() {
		for {
			for index, connection := range h.Connections {
				if err := connection.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
					connection.Conn.Close()
					h.DisconnectUser(index)
					log.Printf("%s disconnect\n", connection.Username)
				}
			}
			time.Sleep(1 * time.Second)
		}
	}()
}
