package http

import "net/http"

// HomeHandler is handler function of home endpoint
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	encodeResponse(w, "Home")
}

// PostMessageHandler is handler function of post message endpoint
func PostMessageHandler(w http.ResponseWriter, r *http.Request) {
	encodeResponse(w, "POST Message")
}

// GetMessageHandler is handler function of get message endpoint
func GetMessageHandler(w http.ResponseWriter, r *http.Request) {
	encodeResponse(w, "GET Message")
}
