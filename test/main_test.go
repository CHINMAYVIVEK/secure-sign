package main

import (
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestAll(t *testing.T) {
	t.Run("User", func(t *testing.T) {
		t.Run("LoginHandler", TestLoginHandler)
	})
}
