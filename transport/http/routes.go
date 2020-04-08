package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/rendyfebry/simple-messaging-api/services"
)

type (
	// PostMessageRequest ...
	PostMessageRequest struct {
		Body string `json:"body"`
	}
)

// MakeHomeHandler ...
func MakeHomeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/index.html")
	})
}

// MakePostMessageHandler ...
func MakePostMessageHandler(svc services.MsgService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		msg, err := svc.CreateMessage(req.Body)
		if err != nil {
			payload := &ErrorResponse{
				Code:    "mssg-500-1",
				Message: err.Error(),
			}

			encodeError(w, payload, http.StatusBadRequest)
			return
		}

		// Broadcast message to
		err = svc.BroadcastMessage(1, msg.Body)
		if err != nil {
			// Do not fail, just print the error
			fmt.Println(err)
		}

		// Return response
		encodeResponse(w, msg)

	})
}

// MakeGetMessageHandler ...
func MakeGetMessageHandler(svc services.MsgService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msgs, err := svc.GetMessages()
		if err != nil {
			payload := &ErrorResponse{
				Code:    "mssg-500-2",
				Message: err.Error(),
			}

			encodeError(w, payload, http.StatusBadRequest)
			return
		}

		encodeResponse(w, msgs)
	})
}

// MakeWebSocketHandler ...
func MakeWebSocketHandler(svc services.MsgService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		svc.RegisterChannel(newConn)

		for {
			// Read message from browser
			msgType, msgBody, err := newConn.ReadMessage()
			if err != nil {
				// Do not fail, just print the error
				fmt.Println("Unable to read message")
				fmt.Println(err)
			}

			// Save message to storage
			msg, err := svc.CreateMessage(string(msgBody))
			if err != nil {
				// Do not fail, just print the error
				fmt.Println("Unable to save message")
				fmt.Println(err)
			}

			// Broadcast to all channels
			err = svc.BroadcastMessage(msgType, msg.Body)
			if err != nil {
				// Do not fail, just print the error
				fmt.Println(err)
			}
		}
	})
}
