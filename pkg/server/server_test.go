package server

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	s := NewServer()

	req, err := http.NewRequest("GET", "/healthcheck", nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()

	s.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "ok", recorder.Body.String())
}
