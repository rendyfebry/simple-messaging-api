package http

import (
	"net/http"

	"github.com/rendyfebry/simple-messaging-api/services"
)

var svc services.MsgService

// PostMessageRequest ...
type PostMessageRequest struct {
	Body string `json:"body"`
}

func init() {
	svc = services.NewService("local")
}

// HomeHandler is handler function of home endpoint
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	encodeResponse(w, "Home")
}

// PostMessageHandler is handler function of post message endpoint
func PostMessageHandler(w http.ResponseWriter, r *http.Request) {
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

	encodeResponse(w, msg)
}

// GetMessageHandler is handler function of get message endpoint
func GetMessageHandler(w http.ResponseWriter, r *http.Request) {
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
}
