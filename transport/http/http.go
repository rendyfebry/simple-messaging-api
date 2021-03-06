package http

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/json-iterator/go"

	"github.com/rendyfebry/simple-messaging-api/services"
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
func MakeRoutes(svc services.MsgService) *mux.Router {
	r := mux.NewRouter()
	r.Handle("/", MakeHomeHandler()).Methods("GET")
	r.Handle("/messages", handlers.LoggingHandler(os.Stdout, MakePostMessageHandler(svc))).Methods("POST")
	r.Handle("/messages", handlers.LoggingHandler(os.Stdout, MakeGetMessageHandler(svc))).Methods("GET")
	r.Handle("/ws", MakeWebSocketHandler(svc))

	r.NotFoundHandler = HandleNotFound()

	return r
}

// HandleNotFound is handler function for all not found endpoint
func HandleNotFound() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload := &ErrorResponse{
			Code:    "mssg-401-1",
			Message: http.StatusText(http.StatusNotFound),
		}

		encodeError(w, payload, http.StatusNotFound)
	})
}

// encodeResponse is http helper function to generate success response
func encodeResponse(w http.ResponseWriter, res interface{}) {
	payload := &SuccessResponse{
		Data: res,
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(payload)
}

// encodeResponse is http helper function to generate error response
func encodeError(w http.ResponseWriter, err *ErrorResponse, errStatus int) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(errStatus)
	json.NewEncoder(w).Encode(err)
}
