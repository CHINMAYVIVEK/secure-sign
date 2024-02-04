package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"secure-sign/app/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

const serverUrl = "http://localhost:8081"

func TestLoginHandler(t *testing.T) {
	apiEndPoint := serverUrl + "/login"
	tests := []struct {
		name       string
		request    *http.Request
		wantStatus int
	}{
		{
			name: "Valid Login",
			request: httptest.NewRequest(
				http.MethodPost,
				apiEndPoint,
				bytes.NewBufferString(`{
					"username": "0123456789",
					"password": "test1234"
				}`),
			),
			wantStatus: http.StatusOK,
		},
		{
			name: "Invalid Login",
			request: httptest.NewRequest(
				http.MethodPost,
				apiEndPoint,
				bytes.NewBufferString(`{
					"username": "username",
					"password": "password"
				}`),
			),
			wantStatus: http.StatusBadRequest,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			user.LoginHandler(recorder, tt.request)

			assert.Equal(t, tt.wantStatus, recorder.Code, "Expected status %d, got %d", tt.wantStatus, recorder.Code)
		})
	}
}
