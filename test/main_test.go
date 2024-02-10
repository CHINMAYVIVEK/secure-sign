package main

import (
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestAll(t *testing.T) {
	t.Run("LoginHandler", TestLoginHandler)
	// t.Run("RegisterHandler", TestRegisterHandler)
	t.Run("GetUserHandler", TestGetUserHandler)
}
