package services

import (
	"time"

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
}

// MsgSvc ...
type MsgSvc struct {
	Database []*Message
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
