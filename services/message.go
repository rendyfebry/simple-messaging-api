package services

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

// Message ...
type Message struct {
	ID        uuid.UUID `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

// MsgService ...
type MsgService interface {
	CreateMessage(string) (*Message, error)
	GetMessages() ([]*Message, error)
	RegisterChannel(newConn *websocket.Conn)
	BroadcastMessage(msgType int, msg string)
}

// MsgSvc ...
type MsgSvc struct {
	Database []*Message
	Channels []*websocket.Conn
}

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

// CreateMessage ...
func (svc *MsgSvc) CreateMessage(body string) (*Message, error) {
	newMsg := &Message{
		ID:        uuid.NewV4(),
		Body:      body,
		CreatedAt: time.Now(),
	}

	svc.Database = append(svc.Database, newMsg)
	return newMsg, nil
}

// GetMessages ...
func (svc *MsgSvc) GetMessages() ([]*Message, error) {
	return svc.Database, nil
}

// RegisterChannel ...
func (svc *MsgSvc) RegisterChannel(newConn *websocket.Conn) {
	svc.Channels = append(svc.Channels, newConn)
}

// BroadcastMessage ...
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
