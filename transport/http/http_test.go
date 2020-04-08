package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeResponse(t *testing.T) {
	w := httptest.NewRecorder()

	encodeResponse(w, "foo")

	ct := w.Header().Get("Content-Type")
	assert.Equal(t, "application/vnd.api+json", ct)
	assert.Equal(t, "{\"data\":\"foo\"}\n", w.Body.String())
}

func TestEncodeError(t *testing.T) {
	w := httptest.NewRecorder()

	var err = &ErrorResponse{
		Code:    "11",
		Message: "SomeError",
	}

	encodeError(w, err, http.StatusBadRequest)

	ct := w.Header().Get("Content-Type")

	assert.Equal(t, "application/vnd.api+json", ct)
	assert.Equal(t, "{\"code\":\"11\",\"message\":\"SomeError\"}\n", w.Body.String())
}
