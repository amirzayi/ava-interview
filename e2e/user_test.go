package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserList(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/user")
	require.NoError(t, err)
	defer resp.Body.Close()
}

func TestCreateUser(t *testing.T) {
	resp, err := http.Post("http://localhost:8080/user", "application/json", http.NoBody)
	require.NoError(t, err)
	defer resp.Body.Close()
}

func TestGetUserByID(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/user/1")
	require.NoError(t, err)
	defer resp.Body.Close()
}

func TestDeleteUserByID(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/user/1")
	require.NoError(t, err)
	defer resp.Body.Close()
}
