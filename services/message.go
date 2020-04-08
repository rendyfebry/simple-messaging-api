package services

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
	uuid "github.com/satori/go.uuid"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

type (
	// Message object
	Message struct {
		ID        uuid.UUID `json:"id"`
		Body      string    `json:"body"`
		CreatedAt time.Time `json:"created_at"`
	}

	// MsgSvc object
	MsgSvc struct {
		Database []*Message
		Channels []*websocket.Conn
	}

	// MsgService interface
	MsgService interface {
		CreateMessage(string) (*Message, error)
		GetMessages() ([]*Message, error)
		RegisterChannel(newConn *websocket.Conn)
		BroadcastMessage(msgType int, msg string)
	}
)

// NewService will create instance of MsgSvc
func NewService(storage string) MsgService {
	// Initialize local storage
	localDB := []*Message{}

	// In the future, if we have other storage,
	// we can implement them here. eg
	// if storage == "sql" {
	//   // call database
	// }

	return &MsgSvc{
		Database: localDB,
		Channels: make([]*websocket.Conn, 0),
	}
}

// CreateMessage will create new message object and save it to db
func (svc *MsgSvc) CreateMessage(body string) (*Message, error) {
	newMsg := &Message{
		ID:        uuid.NewV4(),
		Body:      body,
		CreatedAt: time.Now(),
	}

	svc.Database = append(svc.Database, newMsg)
	return newMsg, nil
}

// GetMessages will return all messages from db
func (svc *MsgSvc) GetMessages() ([]*Message, error) {
	return svc.Database, nil
}

// RegisterChannel will register new connection to the channel list
func (svc *MsgSvc) RegisterChannel(newConn *websocket.Conn) {
	svc.Channels = append(svc.Channels, newConn)
}

// BroadcastMessage will send message to all available channels
func (svc *MsgSvc) BroadcastMessage(msgType int, msg string) {
	msgByte, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, conn := range svc.Channels {
		// Send message to every connection
		if err := conn.WriteMessage(msgType, msgByte); err != nil {
			fmt.Println(err)
			return
		}
	}
}
