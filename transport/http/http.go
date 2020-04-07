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
	// SuccessResponse object
	SuccessResponse struct {
		Data interface{} `json:"data"`
		Meta meta        `json:"meta,omitempty"`
	}

	// ErrorResponse object
	ErrorResponse struct {
		Code    string                 `json:"code"`
		Message string                 `json:"message"`
		Meta    map[string]interface{} `json:"meta,omitempty"`
	}

	meta map[string]interface{}
)

// MakeRoutes will return mux router object
func MakeRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/messages", PostMessageHandler).Methods("POST")
	r.HandleFunc("/messages", GetMessageHandler).Methods("GET")
	r.NotFoundHandler = HandleNotFound()

	return r
}

func encodeResponse(w http.ResponseWriter, res interface{}) {
	payload := &SuccessResponse{
		Data: res,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}

func encodeError(w http.ResponseWriter, err *ErrorResponse, errStatus int) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(errStatus)
	json.NewEncoder(w).Encode(err)
}

// HandleNotFound is handler function for all not found endpoint
func HandleNotFound() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := &ErrorResponse{
			Code:    "http-404",
			Message: http.StatusText(http.StatusNotFound),
		}

		encodeError(w, payload, http.StatusNotFound)
	})
}