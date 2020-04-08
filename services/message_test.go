package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var svc MsgService

func init() {
	svc = NewService("local")
}

func TestNewService(t *testing.T) {
	newSvc := NewService("local")

	assert.NotEmpty(t, newSvc)
	assert.NotNil(t, newSvc)
}

func TestCreateAndGetMessage(t *testing.T) {
	msg, err := svc.CreateMessage("test")
	assert.NoError(t, err)
	assert.NotNil(t, msg)

	mssgs, err := svc.GetMessages()
	assert.NoError(t, err)
	assert.NotEmpty(t, mssgs)
	assert.Equal(t, 1, len(mssgs))
}
