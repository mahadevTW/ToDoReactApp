package handlers

import (
	"net/http"
	"testing"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"encoding/json"
	"github.com/gorilla/csrf"
)

func TestCSRFHandleSuccess(t *testing.T) {
	r, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()

	csrfKey := []byte("csrfkey")
	csrfProtection := csrf.Protect(
		csrfKey,
		csrf.RequestHeader("X-XSRF-Token"),
	)

	csrfProtection(CSRFHandler()).ServeHTTP(w,r)
	csrfToken := &token{}
	err := json.Unmarshal(w.Body.Bytes(), &csrfToken)

	assert.NoError(t,err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t,"application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.NotEmpty(t, csrfToken.CSRFToken)

}