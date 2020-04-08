package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/rendyfebry/simple-messaging-api/services"
)

// MessageRoute ...
type MessageRoute struct {
	svc        services.MsgService
	wsChannels []*websocket.Conn
}

// PostMessageRequest ...
type PostMessageRequest struct {
	Body string `json:"body"`
}

// HomeHandler is handler function of home endpoint
func (mr *MessageRoute) HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/index.html")
	// encodeResponse(w, "Home")
}

// PostMessageHandler is handler function of post message endpoint
func (mr *MessageRoute) PostMessageHandler(w http.ResponseWriter, r *http.Request) {
	var req PostMessageRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		payload := &ErrorResponse{
			Code:    "mssg-400-1",
			Message: err.Error(),
		}

		encodeError(w, payload, http.StatusBadRequest)
		return
	}

	msg, err := mr.svc.CreateMessage(req.Body)
	if err != nil {
		payload := &ErrorResponse{
			Code:    "mssg-500-1",
			Message: err.Error(),
		}

		encodeError(w, payload, http.StatusBadRequest)
		return
	}

	msgByte, err := json.Marshal(msg.Body)
	if err == nil {
		broadcastMsg(mr.wsChannels, 1, msgByte)
	} else {
		fmt.Println(err)
	}

	encodeResponse(w, msg)
}

// GetMessageHandler is handler function of get message endpoint
func (mr *MessageRoute) GetMessageHandler(w http.ResponseWriter, r *http.Request) {
	msgs, err := mr.svc.GetMessages()
	if err != nil {
		payload := &ErrorResponse{
			Code:    "mssg-500-2",
			Message: err.Error(),
		}

		encodeError(w, payload, http.StatusBadRequest)
		return
	}

	encodeResponse(w, msgs)
}

// WebSocketHandler is handler function of get websocket endpoint
func (mr *MessageRoute) WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	newConn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		payload := &ErrorResponse{
			Code:    "mssg-500-3",
			Message: err.Error(),
		}

		encodeError(w, payload, http.StatusBadRequest)
		return
	}

	mr.wsChannels = append(mr.wsChannels, newConn)

	for {
		// Read message from browser
		msgType, msg, err := newConn.ReadMessage()
		if err != nil {
			fmt.Println("Unable to reead message")
			return
		}

		// Save message to storage
		_, err = mr.svc.CreateMessage(string(msg))
		if err != nil {
			fmt.Println("Unable to save message")
			return
		}

		// Broadcast to all channels
		broadcastMsg(mr.wsChannels, msgType, msg)
	}
}
