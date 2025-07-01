package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/amirzayi/ava-interview/api"
	"github.com/stretchr/testify/require"
)

func TestListUser(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/user")
	require.NoError(t, err)
	defer resp.Body.Close()
	require.NotEqual(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestCreateUserShoulFail(t *testing.T) {
	resp, err := http.Post("http://localhost:8080/user", "application/json", http.NoBody)
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestCreateUser(t *testing.T) {
	params := api.User{
		Name:  "amir",
		Phone: "0910",
	}
	b, err := json.Marshal(params)
	require.NoError(t, err)

	resp, err := http.Post("http://localhost:8080/user", "application/json", bytes.NewReader(b))
	require.NoError(t, err)
	defer resp.Body.Close()

	user := api.User{}
	err = json.NewDecoder(resp.Body).Decode(&user)
	require.NoError(t, err)
	require.NotEmpty(t, params)
	require.Equal(t, http.StatusCreated, resp.StatusCode)
	require.Equal(t, params.Name, user.Name)
	require.Equal(t, params.Phone, user.Phone)
}

func TestGetUserByID(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/user/1")
	require.NoError(t, err)

	defer resp.Body.Close()
	user := api.User{}
	err = json.NewDecoder(resp.Body).Decode(&user)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestDeleteUserByID(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/user/1", http.NoBody)
	require.NoError(t, err)
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusNoContent, resp.StatusCode)
}
