package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessfulResponseFromStatusCheckHandler(t *testing.T) {
	r, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	MeAliveMethod(w, r)
	response := w.Body.Bytes()
	assert.Equal(t, "Alive", string(response))
}
