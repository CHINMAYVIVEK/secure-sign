package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"secure-sign/app/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

const serverUrl = "http://localhost:8081/v1/user"

// ResponseRecorder is a custom implementation of http.ResponseWriter for testing purposes.
type ResponseRecorder struct {
	Code       int
	RespHeader http.Header
	Body       *bytes.Buffer
}

func (rw *ResponseRecorder) Header() http.Header {
	return rw.RespHeader
}

func (rw *ResponseRecorder) WriteHeader(statusCode int) {
	rw.Code = statusCode
}

func (rw *ResponseRecorder) Write(data []byte) (int, error) {
	return rw.Body.Write(data)
}

func TestLoginHandler(t *testing.T) {

	apiEndPoint := serverUrl + "/login"
	tests := []struct {
		name       string
		request    *http.Request
		wantStatus int
	}{
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

			assert.Equal(t, tt.wantStatus, recorder.Code, fmt.Sprintf("Expected status %d, got %d", tt.wantStatus, recorder.Code))
			fmt.Printf("Test case %q completed: Expected status %d, got %d\n", tt.name, tt.wantStatus, recorder.Code)
		})
	}
}

// func TestRegisterHandler(t *testing.T) {
// 	apiEndPoint := serverUrl + "/register"
// 	tests := []struct {
// 		name       string
// 		request    *http.Request
// 		wantStatus int
// 	}{
// 		{
// 			name: "Valid Registration",
// 			request: httptest.NewRequest(
// 				http.MethodPost,
// 				apiEndPoint,
// 				bytes.NewBufferString(`{
// 					"first_name": "John",
// 					"last_name": "Doe",
// 					"gender": "Male",
// 					"age": 30,
// 					"email": "john.doe2@example.com",
// 					"phone_number": "9234567892",
// 					"salary": 50000,
// 					"password": "1234"
// 				}`),
// 			),
// 			wantStatus: http.StatusOK,
// 		},
// 		// Add more test cases as needed
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			recorder := httptest.NewRecorder()
// 			user.RegisterHandler(recorder, tt.request)

// 			assert.Equal(t, tt.wantStatus, recorder.Code, fmt.Sprintf("Expected status %d, got %d", tt.wantStatus, recorder.Code))
// 			fmt.Printf("Test case %q completed: Expected status %d, got %d\n", tt.name, tt.wantStatus, recorder.Code)
// 		})
// 	}
// }

func TestGetUserHandler(t *testing.T) {
	apiEndPoint := serverUrl + "9234567891"
	tests := []struct {
		name       string
		request    *http.Request
		wantStatus int
	}{
		{
			name: "Valid User Retrieval",
			request: httptest.NewRequest(
				http.MethodGet,
				apiEndPoint,
				nil,
			),
			wantStatus: http.StatusOK,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			user.GetUserHandler(recorder, tt.request)

			assert.Equal(t, tt.wantStatus, recorder.Code, fmt.Sprintf("Expected status %d, got %d", tt.wantStatus, recorder.Code))
			fmt.Printf("Test case %q completed: Expected status %d, got %d\n", tt.name, tt.wantStatus, recorder.Code)
		})
	}
}
