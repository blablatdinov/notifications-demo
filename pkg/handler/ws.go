package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func (h *Handler) wsGinHandler(c *gin.Context) {
	username, err := getUsername(c)
	if err != nil {
		c.AbortWithStatusJSON(400, ErrorResponse{err.Error()})
		return
	}
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)

	h.services.Notifications.ConnectUser(*conn, username)
	if err != nil {
		return
	}

	for {
		messageChan := h.services.GetChan()
		select {
		case message := <-messageChan:
			err := conn.WriteMessage(1, []byte(message))
			if err != nil {
				log.Fatal(err.Error())
				conn.Close()
			}
		default:
			continue
		}
	}
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}
