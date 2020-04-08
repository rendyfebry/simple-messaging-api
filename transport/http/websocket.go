package http

import (
	"fmt"

	"github.com/gorilla/websocket"
)

var connections []*websocket.Conn

func broadcastMsg(channels []*websocket.Conn, msgType int, msg []byte) {
	for _, conn := range channels {
		// Send message to every connection
		if err := conn.WriteMessage(msgType, msg); err != nil {
			fmt.Println(err)
			return
		}
	}
}
