package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/rendyfebry/simple-messaging-api/services"
)

// MessageRoute ...
type MessageRoute struct {
	svc services.MsgService
}

// PostMessageRequest ...
type PostMessageRequest struct {
	Body string `json:"body"`
}

// HomeHandler is handler function of home endpoint
func (mr *MessageRoute) HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./public/index.html")
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

	// Broadcast message to
	mr.svc.BroadcastMessage(1, msg.Body)

	// Return response
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
	// Create new connect
	newConn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		payload := &ErrorResponse{
			Code:    "mssg-500-3",
			Message: err.Error(),
		}

		encodeError(w, payload, http.StatusBadRequest)
		return
	}

	// Registed conn to channel list
	mr.svc.RegisterChannel(newConn)

	for {
		// Read message from browser
		msgType, msgBody, err := newConn.ReadMessage()
		if err != nil {
			fmt.Println("Unable to read message")
			return
		}

		// Save message to storage
		msg, err := mr.svc.CreateMessage(string(msgBody))
		if err != nil {
			fmt.Println("Unable to save message")
			return
		}

		// Broadcast to all channels
		mr.svc.BroadcastMessage(msgType, msg.Body)
	}
}
