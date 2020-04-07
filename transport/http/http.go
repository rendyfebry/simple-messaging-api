package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/json-iterator/go"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

type (
	// SuccessResponse ...
	SuccessResponse struct {
		Data interface{} `json:"data"`
		Meta meta        `json:"meta,omitempty"`
	}

	meta map[string]interface{}
)

// MakeRoutes will return mux router object
func MakeRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/messages", PostMessageHandler).Methods("POST")
	r.HandleFunc("/messages", GetMessageHandler).Methods("GET")

	return r
}

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

func encodeResponse(w http.ResponseWriter, res interface{}) {
	payload := &SuccessResponse{
		Data: "GET Message",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}
